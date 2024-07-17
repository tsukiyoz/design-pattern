package readwritelock

import (
	"fmt"
	"sync"
	"testing"
)

func TestReadWriteLock(t *testing.T) {
	data := Data{value: 0}

	var wg sync.WaitGroup

	// 多个 goroutine 进行读操作
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Read data:", data.Read())
		}()
	}

	// 单个 goroutine 进行写操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		data.Write(10)
	}()

	// 等待所有 goroutine 完成
	wg.Wait()
}
