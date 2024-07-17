package main

import (
	"fmt"
	"time"
)

func producer(queue chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Producing data:", i)
		queue <- i
	}
	close(queue)
}

func consumer(queue chan int) {
	for data := range queue {
		fmt.Println("Consuming data:", data)
	}
}

func main() {
	queue := make(chan int)

	go producer(queue)
	go consumer(queue)

	time.Sleep(6 * time.Second)
}
