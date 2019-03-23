package grabGeek

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type lessonList struct {
	List []lessonInfo `json:"list"`
}

type articleInfo struct {
	Data lessonList `json:"data"`
	Code int        `json:"code"`
}

type lessonInfo struct {
	Title         string    `json:"article_title"`
	VideoMediaMap videoInfo `json:"video_media_map"`
}

type urlInfo struct {
	Url  string
	Size uint64
}

type videoInfo struct {
	Sd urlInfo `json:"sd"`
	Ld urlInfo `json:"ld"`
	Hd urlInfo `json:"hd"`
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

func getArticleInfoByCourseId(courseId uint64) articleInfo {
	client := &http.Client{
		Transport: &http.Transport{},
	}

	body := strings.NewReader(fmt.Sprintf("{\"cid\":\"%d\",\"size\":200,\"prev\":0,\"orde\":\"earliest\",\"sample\":true}", courseId))

	req, err := http.NewRequest("GET", "https://time.geekbang.org/serv/v1/column/articles", body)
	if err != nil {
		log.Fatal("构建请求失败:" + err.Error())
	}

	req.Header.Set("Cookie", "GCID=5262ddc-9427115-af0561a-011501c; GCESS=BAUEAAAAAAQEgFEBAAMEiwGWXAkBAQsCBAACBIsBllwBBBjsDwAMAQEGBPoUWlwHBDO2rI8KBAAAAAAIAQM-;")
	req.Header.Add("Content-Type", `application/json`)
	referer := fmt.Sprintf(`https://time.geekbang.org/course/intro/%d`, courseId)
	req.Header.Add("Referer", referer)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("视频列表请求失败:" + err.Error())
	}

	byteBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("视频列表内容读取失败:" + err.Error())
	}

	// @todo
	// byteBody = []byte(`{"error":[],"extra":[],"data":{"list":[{"video_media":"{\"sd\":{\"url\":\"https:\\\/\\\/media001.geekbang.org\\\/1e3a867b2592424e9d402a7413d46220\\\/6ea405cfa8314a418389be58e8511a9a-45cb305a5dd4734ebb58e3b5bce730d0-sd.m3u8\",\"size\":91626500},\"ld\":{\"url\":\"https:\\\/\\\/media001.geekbang.org\\\/1e3a867b2592424e9d402a7413d46220\\\/6ea405cfa8314a418389be58e8511a9a-fd2a2d715d7482c28cd1528d5f8c3a7b-ld.m3u8\",\"size\":56695912},\"hd\":{\"url\":\"https:\\\/\\\/media001.geekbang.org\\\/1e3a867b2592424e9d402a7413d46220\\\/6ea405cfa8314a418389be58e8511a9a-f882ae8b20fb62d562e2f0ab8f8a9e37-hd.m3u8\",\"size\":175404000}}","article_subtitle":"","video_cover":"","id":84335,"had_viewed":false,"article_title":"01 | Go语言课程介绍","article_cover":"","video_media_map":{"sd":{"url":"https:\/\/media001.geekbang.org\/1e3a867b2592424e9d402a7413d46220\/6ea405cfa8314a418389be58e8511a9a-45cb305a5dd4734ebb58e3b5bce730d0-sd.m3u8","size":91626500},"ld":{"url":"https:\/\/media001.geekbang.org\/1e3a867b2592424e9d402a7413d46220\/6ea405cfa8314a418389be58e8511a9a-fd2a2d715d7482c28cd1528d5f8c3a7b-ld.m3u8","size":56695912},"hd":{"url":"https:\/\/media001.geekbang.org\/1e3a867b2592424e9d402a7413d46220\/6ea405cfa8314a418389be58e8511a9a-f882ae8b20fb62d562e2f0ab8f8a9e37-hd.m3u8","size":175404000}},"article_could_preview":true,"chapter_id":"207","video_size":130013213,"video_time_arr":{"m":"06","s":"57","h":"00"},"article_summary":"","video_time":"00:06:57","score":1551690060347,"article_ctime":1551690060}]},"code":0}`)

	var articleInfo articleInfo

	if err = json.Unmarshal(byteBody, &articleInfo); err != nil {
		log.Fatal("视频内容json解析错误:" + err.Error())
	}

	return articleInfo
}

func Grab(courseId uint64) {
	articleInfo := getArticleInfoByCourseId(courseId)

	var lessonWaitGroup sync.WaitGroup

	basePath := fmt.Sprintf("%s/%d", basePath, courseId)

	for _, lesson := range articleInfo.Data.List {
		lessonWaitGroup.Add(1)

		go grabLessonVideo(lesson, basePath, &lessonWaitGroup)
	}

	lessonWaitGroup.Wait()
}

// 抓取单课表视频
func grabLessonVideo(info lessonInfo, basePath string, lessonWaitGroup *sync.WaitGroup) {
	defer lessonWaitGroup.Done()

	// 课表根目录
	tmpPath := basePath + "/" + info.Title
	err := os.MkdirAll(tmpPath, 0755)
	if err != nil {
		log.Fatal("创建临时目录:" + err.Error())
	}

	// 获取m3u8文件
	res, err := http.Get(info.VideoMediaMap.Hd.Url)
	if err != nil {
		log.Fatal("获取m3u8文件错误:" + err.Error())
	}

	videoInfoRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("m3u8内容读取错误:" + err.Error())
	}

	m3u8File := tmpPath + "/" + info.VideoMediaMap.Hd.Url[(strings.LastIndex(info.VideoMediaMap.Hd.Url, "/") + 1):]
	if err := ioutil.WriteFile(m3u8File, videoInfoRes, 0755); err != nil {
		log.Fatal("m3u8文件保存失败" + err.Error())
	}

	// 视频基本URL
	videoBaseUrl := info.VideoMediaMap.Hd.Url[:strings.LastIndex(info.VideoMediaMap.Hd.Url, "/")]

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

		go grabTsFile(videoBaseUrl+"/"+videoSuffix, &wg, tmpPath)
	}

	wg.Wait()

	fmt.Println("下载完成" + info.Title)
}

func mergeVideo(m3u8File string, file string, title string) {
	cmd := exec.Command("ffmpeg", "-allowed_extensions", `ALL`, "-i", m3u8File, "-c", `copy`, file+".mp4")
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
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("http内容抓取错误:" + err.Error())
	}

	ts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("http内容读取:" + err.Error())
	}

	// save content into file
	fn := filepath.Join(basePath, url[strings.LastIndex(url, "/"):])
	if err := ioutil.WriteFile(fn, ts, 0666); err != nil {
		log.Fatal("保存视频文件错误:" + err.Error())
	}

	fmt.Println("保存成功: " + fn)
}
