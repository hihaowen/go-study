package _interface

import (
	"fmt"
	"testing"
)

type sendSms interface {
	send(mobile uint64, content string)
}

type gateway string

func (gw gateway) send(mobile uint64, content string) {
	fmt.Printf("gateway: %v send mobile: %d, content: %s\n", gw, mobile, content)
}

func TestInterface(t *testing.T) {
	var aliyunSms sendSms = gateway("aliyun")
	aliyunSms.send(13888888888, "sms 1")

	var zhizhenSms sendSms = gateway("zhizhen")
	zhizhenSms.send(13888888889, "sms 2")
}
