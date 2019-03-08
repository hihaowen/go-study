package _interface

import "fmt"

func Show()  {
	// 空接口
	var I interface{}
	I = "i am str1"

	fmt.Println(I)

	// 获取空接口里的具体原始类型值
	str1, isSet := I.(string)
	fmt.Println(str1, isSet)

	var I2 interface{}
	if I2 == nil {
		fmt.Println("interface is nil")
	}

	// 接口具体应用 (短信发送)
	var sms ISMS // 定义接口
	smsService := SmsService{5, 0}
	sms = &smsService // 这里必须显式的获取指针，否则是不符合接口约束的
	sms.send(138888888, "Bob, how are u?")
	sms.send(138888887, "George, how are u?")
	fmt.Printf("sms send total: %d, success: %d\n", smsService.total, smsService.success)

	// 验证接口类型，类似于PHP中instanceof
	checkSmsService(sms)
}

// 验证接口类型
func checkSmsService(i interface{}) {
	switch t := i.(type) {
	case ISMS:
		fmt.Println("是邮箱服务", t)
	default:
		fmt.Println("是未知服务", t)
	}
}

type ISMS interface {
	send(mobile uint64, content string)
}

type SmsService struct {
	total int
	success int
}

func (sms *SmsService) send(mobile uint64, content string) {
	fmt.Printf("send ok, mobile: %d, content: %s\n", mobile, content)

	// 计数
	sms.success += 1
}
