package grabGeek

import (
	"fmt"
	"testing"
)

func TestGrab(t *testing.T) {
	var courseId uint64

	courseId = 160

	Grab(courseId)

	fmt.Println("全部下载完成")
}

// 测试视频合并
func TestMergeVideo(t *testing.T) {
	mergeVideo("/private/var/folders/0b/tvwm8jns4yvbhj_r5lfhnw3r0000gn/T/160/01 | Go语言课程介绍/6ea405cfa8314a418389be58e8511a9a-f882ae8b20fb62d562e2f0ab8f8a9e37-hd.m3u8", "/private/var/folders/0b/tvwm8jns4yvbhj_r5lfhnw3r0000gn/T/160/01 | Go语言课程介绍", "01 | Go语言课程介绍")
}
