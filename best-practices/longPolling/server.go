package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	address := ":8080"
	log.Printf("Listening and serving HTTP on %s\n", address)

	http.HandleFunc("/long-polling", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		contentChan := make(chan string)

		go getContent(contentChan)

		select {
		case <-time.After(time.Second * 30):
			log.Println("timeout ...")
			fmt.Fprintln(w, "超时了 ...")
			close(contentChan)
		case content := <-contentChan:
			log.Println("content:", content)
			fmt.Fprintln(w, "有数据:", content)
		}

		return
	})

	log.Fatal(http.ListenAndServe(address, nil))

	return
}

func getContent(contentChan chan string) {
	for {
		select {
		case _, ok := <-contentChan:
			if ok == false {
				break
			}
		default:
			content, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/go-study/best-practices/longPolling/txt.dat")
			if err != nil || len(content) == 0 {
				log.Println("数据错误:", err)
				time.Sleep(time.Second * 2)

				continue
			}

			contentChan <- fmt.Sprintf("%s", content)
			break
		}
	}
}
