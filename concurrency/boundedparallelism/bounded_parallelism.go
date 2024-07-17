package boundedparallelism

import (
	"fmt"
	"time"
)

func Worker(id int, semaphore chan struct{}) {
	semaphore <- struct{}{}        // 占用一个信号
	defer func() { <-semaphore }() // 释放信号

	// 模拟任务执行
	fmt.Printf("Worker %d is working\n", id)
	time.Sleep(2 * time.Second)
}
