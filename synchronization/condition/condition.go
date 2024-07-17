package condition

import (
	"fmt"
	"sync"
)

// Data 结构体表示存储数据的对象
type Data struct {
	ready bool       // ready 表示数据是否准备好的标志
	cond  *sync.Cond // cond 用于实现条件变量
}

// NewData 创建一个新的 Data 实例，并初始化 cond 为一个带锁的条件变量
func NewData() *Data {
	return &Data{
		ready: false,
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

// WaitForReady 等待数据准备的方法
func (d *Data) WaitForReady() {
	d.cond.L.Lock()
	defer d.cond.L.Unlock()

	for !d.ready {
		// 等待 cond 的条件变量满足
		d.cond.Wait()
	}

	fmt.Println("Data is ready")
}

// SetReady 设置数据准备好的方法
func (d *Data) SetReady() {
	d.cond.L.Lock()
	defer d.cond.L.Unlock()

	// 设置数据准备好的标志为 true
	d.ready = true
	// 通知等待的线程条件已满足
	d.cond.Signal()
}
