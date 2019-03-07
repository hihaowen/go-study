package _switch

import "fmt"

func Show()  {
	var status = 1
	var statusDesc string

	switch status {
	case getStatusDesc("check"):
		statusDesc = "check"
	case 2:
		statusDesc = "done"
	case 3:
		statusDesc = "delete"
	default:
		statusDesc = "no known"
	}

	fmt.Println("status:", statusDesc)
}

// 根据状态文本获取int类型状态值
func getStatusDesc(status string) int {
	statusMaps := map[string]int{
		"check": 1,
		"done": 2,
		"delete": 3,
 	}

	st, isSet := statusMaps[status]

	if ! isSet {
		return 0
	}

	return st
}
