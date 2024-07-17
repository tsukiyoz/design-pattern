// Package semaphore 实现了 Go 中的信号量弹性模式。
package semaphore

import (
	"errors"
	"time"
)

// ErrNoTickets 是在 Acquire 无法在配置的超时时间内从信号量中获取票证时返回的错误。
var ErrNoTickets = errors.New("无法获取信号量票证")

// Semaphore 实现了信号量弹性模式。
type Semaphore struct {
	sem     chan struct{} // sem 是一个通道，用于表示信号量
	timeout time.Duration // timeout 指定等待获取票证的持续时间
}

// New 使用给定的票证数量和超时时间构造一个新的 Semaphore。
func New(tickets int, timeout time.Duration) *Semaphore {
	return &Semaphore{
		sem:     make(chan struct{}, tickets), // 使用指定数量的票证初始化信号量通道
		timeout: timeout,                      // 设置获取票证的超时时间
	}
}

// Acquire 尝试从信号量中获取一个票证。如果成功，返回 nil。
// 如果在超时时间内无法获取，则返回 ErrNoTickets。在单个 Semaphore 上并发调用 Acquire 是安全的。
func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}: // 尝试发送一个空结构体到信号量通道
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

// Release 将一个已获取的票证释放回信号量。在单个 Semaphore 上并发调用 Release 是安全的。
// 在未先获取票证的 Semaphore 上调用 Release 是错误的。
func (s *Semaphore) Release() {
	<-s.sem
}

// IsEmpty 在某一时刻如果没有任何票证被持有，则返回 true。
// 可以与 Acquire 和 Release 并发调用 IsEmpty，但请注意结果可能不可预测。
func (s *Semaphore) IsEmpty() bool {
	return len(s.sem) == 0
}
