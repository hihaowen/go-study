package channel

import (
	"fmt"
	"testing"
	"time"
)

type Worker struct {
	id int
}

func (worker Worker) doWork(ch chan int) {
	for {
		fmt.Printf("Worker %d is working: %d\n", worker.id, <-ch)
		time.Sleep(time.Second * 5)
	}
}

func TestDefaultChannel(t *testing.T) {
	// 创建信道
	ch1 := make(chan int, 6)

	// 创建几个Worker
	for i := 0; i < 5; i++ {
		worker := Worker{i}
		go worker.doWork(ch1)
	}

	// 分配工作
	for i := 0; i < 10; i++ {
		ch1 <- i
		fmt.Println("通道等待中...", len(ch1))
	}

	time.Sleep(time.Second * 10)

	fmt.Println("主程序执行完了")
}
