package main

import (
	"sync"
)

// Merge 函数实现了 FanIn 操作，将不同的通道组合成一个
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int, 3)

	wg.Add(len(cs))

	// send 函数从输入通道 c 中复制值发送到 out，直到 c 关闭，然后调用 wg.Done
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// 启动多个 goroutine 开始工作
	for _, c := range cs {
		go send(c)
	}

	// 启动一个 goroutine，在所有 send goroutine 完成后关闭 out。必须在 wg.Add 调用之后开始
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// generateNumbersPipeline 函数生成一个通道，用于发送给定的整数切片中的数字
func generateNumbersPipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out) // 发送完成后关闭通道
	}()
	return out
}

// squareNumber 函数接收一个整数通道，并将每个收到的数字平方后发送到新通道中
func squareNumber(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out) // 发送完成后关闭通道
	}()
	return out
}
