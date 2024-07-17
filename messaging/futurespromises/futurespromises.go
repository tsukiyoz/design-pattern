package main

import (
	"fmt"
	"time"
)

func asyncOperation() <-chan int {
	result := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		result <- 42
	}()

	return result
}

func main() {
	future := asyncOperation()

	fmt.Println("Performing other operations...")

	value := <-future
	fmt.Println("Async operation result:", value)
}
