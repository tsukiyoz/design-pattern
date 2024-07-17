package monitor

import (
	"fmt"
	"testing"
)

func TestMonitor(t *testing.T) {
	// 创建一个新的 Monitor 实例
	monitor := NewMonitor()

	// 启动一个协程，调用 WriteData 方法向 Monitor 写入数据
	go func() {
		monitor.WriteData(10)
	}()

	// 等待数据被写入并读取
	data := monitor.ReadData()
	fmt.Println("已读取数据:", data)
}
