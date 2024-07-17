package monitor

import (
	"fmt"
	"sync"
)

// Monitor 结构体定义了 Monitor（监视器）类型
type Monitor struct {
	mu   sync.Mutex // 互斥锁用于保护共享资源
	cond *sync.Cond // 条件变量用于协调协程的同步
	data int        // 共享数据
}

// NewMonitor 创建并返回一个新的 Monitor 实例
func NewMonitor() *Monitor {
	m := &Monitor{}
	m.cond = sync.NewCond(&m.mu) // 使用 Monitor 结构体中的互斥锁创建条件变量
	return m
}

// WriteData 方法用于向 data 写入数据并发送信号通知等待的协程
func (m *Monitor) WriteData(d int) {
	m.mu.Lock()         // 加锁
	defer m.mu.Unlock() // 延迟解锁

	m.data = d
	fmt.Println("已写入数据:", d)
	m.cond.Signal() // 发送信号通知等待的协程
}

// ReadData 方法用于等待 data 被设置后读取数据
func (m *Monitor) ReadData() int {
	m.mu.Lock()         // 加锁
	defer m.mu.Unlock() // 延迟解锁

	for m.data == 0 {
		m.cond.Wait() // 等待条件变量满足
	}
	fmt.Println("已读取数据:", m.data)
	return m.data
}
