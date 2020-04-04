package main

import (
	"bufio"
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
	bp, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("获取当前目录地址" + err.Error())
	}
	return bp + "/v"
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

	partFilesDir := basePath + "/" + *title
	err := os.MkdirAll(partFilesDir, 0755)
	if err != nil {
		log.Fatal("mkdir error:", err.Error())
	}

	// 1.grab m3u8 file、key file
	m3u8File := grabIntoFile(m3u8Url, partFilesDir+"/index.m3u8")
	log.Println("m3u8:", m3u8File)

	keyFile := grabIntoFile(keyUrl, partFilesDir+"/index.key")
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

		videoFile := partFilesDir + ".mp4"

		mergeVideo(m3u8File, videoFile)

		// 3.2 remove part files dir
		err = os.RemoveAll(partFilesDir)
		if err != nil {
			log.Fatalf("remove part files dir error: %s, dir: %s", err.Error(), partFilesDir)
		}
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

		go func(tsUrl, partFilesDir string) {
			runProcessorsChan <- 1
			log.Println("fetch video from:", tsUrl)
			tsFile := grabIntoFile(tsUrl, partFilesDir+"/"+tsUrl[strings.LastIndex(tsUrl, "/"):])
			log.Println("fetch video ok:", tsFile)
			<-runProcessorsChan
			quitProcessorsChan <- true
		}(string(l), partFilesDir)
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
	cmd := exec.Command("ffmpeg", "-allowed_extensions", `ALL`, "-i", m3u8File, "-c", `copy`, toFile)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("merge failed:", toFile, err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("merge failed:", toFile, err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	log.Printf("merge and print: %s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal("merge failed:", toFile, err)
	}

	log.Println("merge success:", toFile)
}

type retryLock struct {
	retry map[string]uint
	lock  sync.RWMutex
}

func (r *retryLock) set(uniq string, retry uint) {
	defer r.lock.Unlock()
	r.lock.Lock()
	r.retry[uniq] = retry
}

func (r *retryLock) get(uniq string) uint {
	defer r.lock.RUnlock()
	r.lock.RLock()
	v, _ := r.retry[uniq]
	return v
}

// 抓取视频重试锁
var grabRetryLock = newRetryLock()

func newRetryLock() *retryLock {
	retry := make(map[string]uint, 0)
	return &retryLock{retry: retry}
}

func grabIntoFile(url, toFile string) string {
	// if file exists
	_, err := os.Stat(toFile)
	if err == nil {
		log.Printf("grab from local cache: %s\n", toFile)
		return toFile
	}

	res, err := httpClient.Get(url)
	if err != nil {
		retry := grabRetryLock.get(url)
		if retry < 30 {
			retry++
			grabRetryLock.set(url, retry)
			log.Printf("http request error: %s, url: %s, retry: %d\n", err.Error(), url, retry)
			return grabIntoFile(url, toFile)
		}
		log.Fatalf("http request error: %s, url: %s\n", err.Error(), url)
	}

	ts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		retry := grabRetryLock.get(url)
		if retry < 30 {
			retry++
			grabRetryLock.set(url, retry)
			log.Printf("http response error: %s, url: %s, retry: %d\n", err.Error(), url, retry)
			return grabIntoFile(url, toFile)
		}
		log.Fatalf("http response error: %s, url: %s\n", err.Error(), url)
	}

	// save content into file
	if err := ioutil.WriteFile(toFile, ts, 0666); err != nil {
		log.Fatal("save file error:", err.Error())
	}

	return toFile
}
