package main

import (
	"errors"
	"fmt"
	"time"
)

// ErrTimedOut 是在截止时间过期时从 Run 返回的错误。
var ErrTimedOut = errors.New("Timed out waiting for function to finish")

// Deadline 实现了截止时间/超时恢复模式。
type Deadline struct {
	timeout time.Duration
}

// New 使用给定的超时时间构造一个新的 Deadline。
func New(timeout time.Duration) *Deadline {
	return &Deadline{
		timeout: timeout,
	}
}

func (d *Deadline) Run(work func(<-chan struct{}) error) error {
	result := make(chan error)     // 用于接收工作函数的结果
	stopper := make(chan struct{}) // 用于停止工作函数的信号

	go func() {
		value := work(stopper) // 调用工作函数
		select {
		case result <- value: // 将工作函数的结果发送到 result 通道
		case <-stopper: // 如果收到停止信号，则不发送结果
		}
	}()

	select {
	case ret := <-result: // 等待工作函数的结果
		return ret
	case <-time.After(d.timeout): // 如果超过截止时间
		close(stopper)     // 关闭 stopper 通道，停止工作函数
		return ErrTimedOut // 返回超时错误
	}
}

func main() {
	dl := New(1 * time.Second) // 创建一个截止时间为 1 秒的 Deadline

	err := dl.Run(func(stopper <-chan struct{}) error {
		time.Sleep(2 * time.Second) // 模拟一个可能耗时的操作
		// 做一些可能耗时的操作
		// 检查 stopper 函数，如果超时则放弃
		fmt.Println("任务成功完成")
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}
