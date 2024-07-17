package barrier

import (
	"fmt"
	"sync"
	"testing"
)

func TestNBarrier(t *testing.T) {
	var barrier sync.WaitGroup
	numWorkers := 3 // 设定任务数量为 3

	// 设置等待的任务数量
	barrier.Add(numWorkers)

	// 启动多个 Goroutines 执行任务
	for i := 0; i < numWorkers; i++ {
		go Worker(i, &barrier)
	}

	// 等待所有任务完成
	barrier.Wait()

	fmt.Println("All workers have finished, continue to the next step")
}
