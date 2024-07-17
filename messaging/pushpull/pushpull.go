package main

import (
	"fmt"
	"time"
)

// Push 模式示例
func pushData(dataCh chan<- int) {
	for i := 0; i < 5; i++ {
		dataCh <- i
		time.Sleep(time.Second)
	}
	close(dataCh)
}

// Pull 模式示例
func pullData(dataCh <-chan int) {
	for data := range dataCh {
		fmt.Println("Pulled data:", data)
	}
}

func main() {
	pushCh := make(chan int)
	pullCh := make(chan int)

	go pushData(pushCh)
	go pullData(pullCh)

	for i := 0; i < 5; i++ {
		value := <-pushCh
		pullCh <- value
	}

	time.Sleep(5 * time.Second)
}
