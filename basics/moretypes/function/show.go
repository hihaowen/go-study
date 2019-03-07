package function

import "fmt"

func Show()  {
	// 闭包
	getStrArrLen := func(strArr []string) int {
		return len(strArr)
	}

	strArr := []string{"a", "b"}

	fmt.Println("闭包函数 - 根据数组算长度: ",getStrArrLen(strArr))

	// 函数作为值
	plus := plus()
	fmt.Println("1 + 2 = ", plus(1, 2))

	// 函数作为参数 (抽象化、将短信内容与发送网关解耦)
	// 	短信内容
	smsMessage := SmsMessage{13888888888, "您的短信验证码是: xxxx"}
	// 	阿里云、蓝汛短信网关
	aliyunSmsGateway := func(mobile int, content string) bool {
		fmt.Println("sms send by aliyun")
		return true
	}
	lanxunSmsGateway := func(mobile int, content string) bool {
		fmt.Println("sms send by lanxun")

		return false
	}
	// 	短信发送
	sendSMS(aliyunSmsGateway, smsMessage)
	sendSMS(lanxunSmsGateway, smsMessage)

	// 闭包函数调用外函数内的参数、方法
	incrVal := incr()
	fmt.Println("闭包内调用外函数内参数后的值:", incrVal())
	fmt.Println("闭包内调用外函数内参数后的值:", incrVal())
	fmt.Println("闭包内调用外函数内参数后的值:", incrVal())
}

// 短信内容
type SmsMessage struct {
	mobile int
	content string
}

// 短信发送
func sendSMS(gateway func(int, string) bool, message SmsMessage) {
	ret := gateway(message.mobile, message.content)

	if ! ret {
		fmt.Println("sms send failed")
		return
	}

	fmt.Println("sms send ok")
}

// 两个数的加法运算
func plus() func(int, int) int {
	return func(a, b int) int {
		return a + b
	}
}

// 闭包函数调用外函数内的参数、方法
func incr() func() int {
	base := 0
	return func() int {
		base ++
		return base
	}
}
