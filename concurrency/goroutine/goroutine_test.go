package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 主进程一定要在子线程之后执行完成
func TestGoroutine(t *testing.T) {
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("我应该是第二个被打印出来的")
	}()

	fmt.Println("我应该是第一个被打印出来的")

	time.Sleep(time.Second * 3)
}

var counter int
var mux sync.Mutex

func TestGoroutineCounter(t *testing.T) {
	for i := 0; i < 10; i++ {
		go incr()
	}

	time.Sleep(time.Second * 2)

	fmt.Println("最终counter=", counter)
}

func incr() {
	// unlock
	defer mux.Unlock()

	// lock
	mux.Lock()
	counter++
	fmt.Println("当前counter=", counter)
}
