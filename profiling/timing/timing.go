package main

import (
	"fmt"
	"time"
)

func task() {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("Task completed in %s\n", elapsed)
	}()

	// 模拟任务执行
	sum := 0
	for i := 1; i <= 1000000; i++ {
		sum += i
	}
}

func main() {
	fmt.Println("Timing Functions Example")
	task()
}
