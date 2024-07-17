package barrier

import (
	"fmt"
	"sync"
)

func Worker(id int, barrier *sync.WaitGroup) {
	defer barrier.Done() // 每个任务完成时标记 Done

	fmt.Printf("Worker %d starts\n", id)
	// 模拟任务执行
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d is working\n", id)
	}
	fmt.Printf("Worker %d finishes\n", id)
}
