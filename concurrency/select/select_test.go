package _select

import (
	"fmt"
	"testing"
	"time"
)

type Worker struct {
	id int
}

func (worker Worker) doWork(workChannel chan int, finishedChannel chan bool) {
	for {
		select {
		case workId := <-workChannel:
			fmt.Printf("Worker %d is working ... %d\n", worker.id, workId)
		case <-time.After(time.Second * 3):
			goto finished
		}
	}
finished:
	fmt.Printf("Worker %d is working timeout\n", worker.id)
	finishedChannel <- true
}

func TestSelectWithGoroutineAndChannel(t *testing.T) {
	workChannel := make(chan int)
	finishedChannel := make(chan bool)

	for i := 0; i < 5; i++ {
		go Worker{i}.doWork(workChannel, finishedChannel)
	}

	for i := 0; i < 10; i++ {
		workChannel <- i
	}

	// time.Sleep(time.Second * 5)

	select {
	case finished := <-finishedChannel:
		if finished == true {
			fmt.Println("all work done.")
		}
	}
}
