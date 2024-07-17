package facade

import "fmt"

// VegVendor 蔬菜供应商
type VegVendor struct{}

func (v *VegVendor) purchase() {
	fmt.Println("供应蔬菜")
}

// Chef 厨师
type Chef struct{}

func (c *Chef) Cook() {
	fmt.Println("下厨烹饪")
}

// Waiter 服务员
type Waiter struct{}

func (h *Waiter) Service() {
	fmt.Println("开始为用户服务")
}

// Cleaner 清洁工
type Cleaner struct{}

func (c *Cleaner) Clean() {
	fmt.Println("开始清洁卫生")
}

// Facade 外观结构体
type Facade struct {
	vendor  *VegVendor
	chef    *Chef
	waiter  *Waiter
	cleaner *Cleaner
}

// NewFacade 创建一个新的外观对象
func NewFacade() *Facade {
	facade := &Facade{}
	facade.cleaner = &Cleaner{}
	facade.chef = &Chef{}
	facade.waiter = &Waiter{}
	facade.cleaner = &Cleaner{}

	return facade
}

// Order 开始处理订单流程
func (f *Facade) Order() {
	// 服务员接待入座
	f.waiter.Service()
	// 厨师开始烹饪
	f.chef.Cook()
	// 上菜
	f.waiter.Service()
	// 收拾桌子洗碗
	f.cleaner.Clean()
}
