package lockmutex

import (
	"sync"
	"testing"
)

// TestLockMutex 是对 Lock/Mutex 模式的测试函数
func TestLockMutex(t *testing.T) {
	// 创建一个新的计数器实例
	counter := NewCounter()

	var wg sync.WaitGroup
	// 启动5个协程并发执行增加计数器操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 调用计数器的增加方法
			counter.Increment()
		}()
	}

	// 等待所有协程完成
	wg.Wait()
}
