package json_and_go

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

type Message struct {
	Name string
	Age  int
}

func TestJsonEncode(t *testing.T) {
	msg := []Message{
		{"George", 18},
		{"Bob", 30},
	}

	b, err := json.Marshal(msg)

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}

func TestJsonDecode(t *testing.T) {
	text := []byte(`[{"Name":"George","Age":18},{"Name":"Bob","Age":30}]`)

	var msgs []Message

	err := json.Unmarshal(text, &msgs)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range msgs {
		fmt.Println(msg.Name, msg.Age)
	}
}
