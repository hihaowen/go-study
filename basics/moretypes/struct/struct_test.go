package _struct

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

// optional string literal tag
// 可选字符串文字标记
// https://www.cnblogs.com/ycyoes/p/8416154.html

func TestOptionStringTag(t *testing.T) {

	type user struct {
		Id     uint32 `json:"id"`
		Name   string `json:"name"`
		Age    uint8  `json:"age"`
		Mobile uint64 `json:"mob,omitempty"`
	}

	u := user{
		1,"George", 18, 13801010101,
	}

	b, err := json.Marshal(u)

	if err != nil {
		log.Fatal(err)
		return
	}

	os.Stdout.Write(b)
}
