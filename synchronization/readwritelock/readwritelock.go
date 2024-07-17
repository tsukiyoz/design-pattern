package readwritelock

import (
	"sync"
)

// Data 结构体定义了包含读写锁和数据值的类型
type Data struct {
	rwLock sync.RWMutex // 读写锁，用于并发访问控制
	value  int          // 数据值
}

// Read 方法用于读取数据值，使用读锁定保证并发安全
func (d *Data) Read() int {
	d.rwLock.RLock()         // 读锁定
	defer d.rwLock.RUnlock() // 延迟释放读锁

	return d.value
}

// Write 方法用于写入数据值，使用写锁定保证并发安全
func (d *Data) Write(val int) {
	d.rwLock.Lock()         // 写锁定
	defer d.rwLock.Unlock() // 延迟释放写锁

	d.value = val
}
