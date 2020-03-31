package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var uuid = flag.String("uuid", "", "UUID")
var title = flag.String("title", "", "Title")

// 基础设置
func init() {
	// 参数解析
	flag.Parse()

	if *uuid == "" || *title == "" {
		log.Fatal("params error")
	}
}

// 当前运行文件目录
var basePath = func() string {
	// 获取当前目录地址
	basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("获取当前目录地址" + err.Error())
	}
	return basePath + "/v"
}()

// http client
var httpClient = &http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 60 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   2 * time.Second,
		ResponseHeaderTimeout: 3 * time.Second,
		ExpectContinueTimeout: 3 * time.Second,
	},
}

func main() {
	m3u8Url := fmt.Sprintf("https://www.zuo9.live/api/app/m3u8/index.m3u8?uuid=%s&token=123456&nonce=123456&rate=720", *uuid)
	keyUrl := fmt.Sprintf("https://www.zuo9.live/api/app/m3u8/index.key?uuid=%s&rate=720", *uuid)

	// toDir := md5Sum(m3u8Url)
	toDir := *title
	err := os.MkdirAll(basePath+"/"+toDir, 0755)
	if err != nil {
		log.Fatal("mkdir error:", err.Error())
	}

	// 1.grab m3u8 file、key file
	m3u8File := grabIntoFile(m3u8Url, toDir+"/index.m3u8")
	log.Println("m3u8:", m3u8File)

	keyFile := grabIntoFile(keyUrl, toDir+"/index.key")
	log.Println("key:", keyFile)

	// 2.read m3u8 file, for fetch per valid line and pull video info, save into a file
	fi, err := os.Open(m3u8File)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return
	}

	defer fi.Close()

	// 3.merge all part video files to a fully file
	defer func() {
		// 3.1 modify m3u8 file
		m3u8FileModify(m3u8File)

		mergeVideo(m3u8File, toDir+"/index.mp4")

		/*
			err = os.RemoveAll(toDir)
			if err != nil {
				log.Fatal("清理临时保存目录:" + err.Error())
			}
		*/
	}()

	// 正在运行的processors channel
	runProcessorsChan := make(chan int, 10) // 保证同时在运行的processors为10个
	// 退出的processors channel
	quitProcessorsChan := make(chan bool)
	// processors总数
	processors := 0

	br := bufio.NewReader(fi)
	for {
		l, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if string(l[0:1]) == "#" {
			continue
		}

		processors++

		go func(tsUrl, toDir string) {
			runProcessorsChan <- 1
			log.Println("fetch video from:", tsUrl)
			tsFile := grabTsFile(tsUrl, toDir+"/"+tsUrl[strings.LastIndex(tsUrl, "/"):])
			log.Println("fetch video ok:", tsFile)
			<-runProcessorsChan
			quitProcessorsChan <- true
		}(string(l), toDir)
	}

	// 等待所有进程退出
	for i := 0; i < processors; i++ {
		<-quitProcessorsChan
	}

	log.Println("all part video files done")
}

func m3u8FileModify(m3u8File string) {
	// sed -i '' "6s/URI\=\"[^\"]*\"/URI\=\"index\.key\"/" index.m3u8
	cmd := exec.Command("sed", "-i", "", "6s/URI\\=\"[^\"]*\"/URI\\=\"index\\.key\"/", m3u8File)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// sed -i '' "s/^https\:\/\/\([^\/]*\/\)\{1,\}\([^\/]*\.ts\)$/\2/" index.m3u8
	cmd = exec.Command("sed", "-i", "", "s/^https\\:\\/\\/\\([^\\/]*\\/\\)\\{1,\\}\\([^\\/]*\\.ts\\)$/\\2/", m3u8File)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func mergeVideo(m3u8File string, toFile string) {
	cmd := exec.Command("ffmpeg", "-allowed_extensions", `ALL`, "-i", m3u8File, "-c", `copy`, basePath+"/"+toFile)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("视频合并失败:", toFile, err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("视频合并失败:", toFile, err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	log.Printf("视频合并打印: %s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal("视频合并失败:", toFile, err)
	}

	log.Println("视频合并成功", toFile)
}

type grabTsRetry struct {
	retry map[string]uint
	lock  sync.RWMutex
}

func (r *grabTsRetry) set(uniq string, retry uint) {
	defer r.lock.Unlock()
	r.lock.Lock()
	r.retry[uniq] = retry
}

func (r *grabTsRetry) get(uniq string) uint {
	defer r.lock.RUnlock()
	r.lock.RLock()
	v, _ := r.retry[uniq]
	return v
}

var grabTsRetryLock = newGrabTsRetryLock()

func newGrabTsRetryLock() *grabTsRetry {
	retry := make(map[string]uint, 0)
	return &grabTsRetry{retry: retry}
}

// 根据地址抓取并下载保存
func grabTsFile(url, toFile string) string {

	res, err := httpClient.Get(url)
	if err != nil {
		retry := grabTsRetryLock.get(url)
		if retry < 3 {
			log.Printf("http request error: %s, retry: %d\n", err.Error(), retry)
			return grabTsFile(url, toFile)
		}

		grabTsRetryLock.set(url, retry+1)
		log.Fatal("http request error:", err.Error())
	}

	ts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		retry := grabTsRetryLock.get(url)
		if retry < 3 {
			log.Printf("http get response error: %s, retry: %d\n", err.Error(), retry)
			return grabTsFile(url, toFile)
		}

		grabTsRetryLock.set(url, retry+1)
		log.Fatal("http get response error:", err.Error())
	}

	// save content into file
	fn := filepath.Join(basePath, toFile)
	if err := ioutil.WriteFile(fn, ts, 0666); err != nil {
		log.Fatal("save ts file error:", err.Error())
	}

	return fn
}

func grabIntoFile(url, toFile string) string {
	// http grab content
	res, err := httpClient.Get(url)
	if err != nil {
		log.Fatal("http内容抓取错误:", err.Error())
	}

	ts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("http内容读取:", err.Error())
	}

	// save content into file
	fn := filepath.Join(basePath, toFile)
	if err := ioutil.WriteFile(fn, ts, 0666); err != nil {
		log.Fatal("保存文件错误:", err.Error())
	}

	return fn
}

func md5Sum(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
