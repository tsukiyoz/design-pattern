package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second) // 模拟耗时操作
		fmt.Printf("Goroutine 1: %d\n", i)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		time.Sleep(time.Second) // 模拟耗时操作
		fmt.Printf("Goroutine 2: %c\n", i)
	}
}

func main() {
	go printNumbers() // 启动第一个 Goroutine
	go printLetters() // 启动第二个 Goroutine

	time.Sleep(4 * time.Second) // 等待 Goroutines 执行完成
	fmt.Println("Main Goroutine exits")
}
