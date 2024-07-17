package lockmutex

import (
	"fmt"
	"sync"
)

// Counter 结构体用于实现一个计数器对象
type Counter struct {
	mu    sync.Mutex // 互斥锁用于保护计数器的并发访问
	count int        // count 记录计数器的当前值
}

// NewCounter 创建一个新的计数器实例
func NewCounter() *Counter {
	return &Counter{}
}

// Increment 增加计数器的方法
func (c *Counter) Increment() {
	c.mu.Lock()         // 加锁，防止并发访问
	defer c.mu.Unlock() // 延迟解锁，确保互斥锁被正确释放

	c.count++ // 增加计数器的值
	fmt.Println("增加后的计数:", c.count)
}
