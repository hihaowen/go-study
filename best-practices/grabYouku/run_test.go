package grabYouku

import (
	"fmt"
	"testing"
)

func TestGrab(t *testing.T) {

	m3u8Url := "https://valipl-vip.cp31.ott.cibntv.net/6572D338DEB4171A85FAF38E8/03000700005D284928B07503AA792EC7E4B2DD-E13C-44B5-B331-28EF56E13EBE-1-114.m3u8?ccode=0502&duration=4703&expire=18000&psid=330c83a7e608d67de2f833deff801467&ups_client_netip=3b6d9926&ups_ts=1564411490&ups_userid=786024432&utid=iI3OEO3M%2FHACAT2HqVAfKPCr&vid=XNDI1MjQwNzk0MA&vkey=A97823d1db761493f126a95bdea24139d&sm=1&operate_type=1&bc=2"

	GrabVideo(VideoInfo{"水猴子", m3u8Url})

	fmt.Println("全部下载完成")
}

// 测试视频合并
func TestMergeVideo(t *testing.T) {
	mergeVideo("/private/var/folders/0b/tvwm8jns4yvbhj_r5lfhnw3r0000gn/T/160/01 | Go语言课程介绍/6ea405cfa8314a418389be58e8511a9a-f882ae8b20fb62d562e2f0ab8f8a9e37-hd.m3u8", "/private/var/folders/0b/tvwm8jns4yvbhj_r5lfhnw3r0000gn/T/160/01 | Go语言课程介绍", "01 | Go语言课程介绍")
}
