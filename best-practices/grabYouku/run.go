package grabYouku

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type VideoInfo struct {
	Title   string
	M3u8Url string
}

// 当前运行文件目录
var basePath = func() string {
	// 获取当前目录地址
	basePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("获取当前目录地址" + err.Error())
	}

	return basePath
}()

// 抓取视频
func GrabVideo(info VideoInfo) {

	// 课表根目录
	tmpPath := basePath + "/" + info.Title
	err := os.MkdirAll(tmpPath, 0755)
	if err != nil {
		log.Fatal("创建临时目录:" + err.Error())
	}

	// 获取m3u8文件
	res, err := http.Get(info.M3u8Url)
	if err != nil {
		log.Fatal("获取m3u8文件错误:" + err.Error())
	}

	videoInfoRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("m3u8内容读取错误:" + err.Error())
	}

	m3u8File := tmpPath + "/" + info.M3u8Url[(strings.LastIndex(info.M3u8Url, "/") + 1):]
	m3u8File = m3u8File[0:(strings.LastIndex(m3u8File, "?"))]
	if err := ioutil.WriteFile(m3u8File, videoInfoRes, 0755); err != nil {
		log.Fatal("m3u8文件保存失败" + err.Error())
	}

	defer func() {
		mergeVideo(m3u8File, tmpPath, info.Title)

		err = os.RemoveAll(tmpPath)
		if err != nil {
			log.Fatal("清理临时保存目录:" + err.Error())
		}
	}()

	// 拉取视频并保存
	var wg sync.WaitGroup

	videoInfoResText := fmt.Sprintf("%s", videoInfoRes)

	for _, videoSuffix := range strings.Split(videoInfoResText, "\n") {
		if strings.HasPrefix(videoSuffix, "#") || videoSuffix == "" {
			continue
		}

		wg.Add(1)

		// go grabTsFile(videoSuffix, &wg, tmpPath)
		// 如果用协程会有超时风险
		grabTsFile(videoSuffix, &wg, tmpPath)
	}

	wg.Wait()

	fmt.Println("下载完成" + info.Title)
}

func mergeVideo(m3u8File string, file string, title string) {
	cmd := exec.Command("ffmpeg", "-allowed_extensions", `ALL`, "-protocol_whitelist", `file,http,https,tcp,tls`, "-i", m3u8File, "-c", `copy`, file+".mp4")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal("视频合并失败:", title, err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("视频合并失败:", title, err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("视频合并打印: %s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal("视频合并失败:", title, err)
	}

	log.Println("视频合并成功", title)
}

// 根据地址抓取并下载保存
func grabTsFile(url string, wg *sync.WaitGroup, basePath string) {
	defer wg.Done()

	// http grab content
	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	res, err := c.Get(url)
	if err != nil {
		log.Fatal("http内容抓取错误:" + err.Error())
	}

	ts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("http内容读取:" + err.Error())
	}

	// save content into file
	fn := filepath.Join(basePath, url[strings.LastIndex(url, "/"):])
	fn = fn[0:strings.LastIndex(fn, "?")]
	if err := ioutil.WriteFile(fn, ts, 0666); err != nil {
		log.Fatal("保存视频文件错误:" + err.Error())
	}

	fmt.Println("保存成功: " + fn)
}
