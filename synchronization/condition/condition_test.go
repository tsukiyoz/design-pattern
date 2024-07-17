package condition

import (
	"fmt"
	"testing"
)

func TestConditionVariable(t *testing.T) {
	data := NewData()

	go func() {
		data.WaitForReady()
	}()

	// 模拟数据准备过程
	fmt.Println("Preparing data...")
	// 模拟数据准备完毕
	data.SetReady()

	// 等待子协程完成
	fmt.Println("Waiting for sub goroutine to finish...")
	// 这里可以做其他操作
}
