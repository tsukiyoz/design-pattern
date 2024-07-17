package decorator

import "fmt"

// Showable 是展示行为的接口
type Showable interface {
	Show() // 展示
}

// Girl 表示女生类，实现了 Showable 接口
type Girl struct{}

// NewGirl 创建一个新的女生实例
func NewGirl() *Girl {
	return &Girl{}
}

// Show 展示女生的素颜
func (g *Girl) Show() {
	fmt.Print("女生的素颜")
}

// Decorator 是装饰器的基类
type Decorator struct {
	showable Showable // 包含 Showable 接口的字段
}

// NewDecorator 创建一个新的装饰器
func NewDecorator(showable Showable) *Decorator {
	return &Decorator{showable: showable}
}

// Show 展示装饰器的行为
func (d *Decorator) Show() {
	d.showable.Show()
}

// FoundationMakeup 表示粉底装饰器
type FoundationMakeup struct {
	Decorator // 继承自 Decorator 基类
}

// NewFoundationMakeup 创建一个新的粉底装饰器
func NewFoundationMakeup(showable Showable) *FoundationMakeup {
	return &FoundationMakeup{Decorator: Decorator{showable: showable}}
}

// Show 展示打粉底的行为
func (f *FoundationMakeup) Show() {
	fmt.Print("打粉底【")
	f.Decorator.Show()
	fmt.Print("】")
}

// Lipstick 表示口红装饰器
type Lipstick struct {
	Decorator // 继承自 Decorator 基类
}

// NewLipstick 创建一个新的口红装饰器
func NewLipstick(showable Showable) *Lipstick {
	return &Lipstick{Decorator: Decorator{showable: showable}}
}

// Show 展示涂口红的行为
func (l *Lipstick) Show() {
	fmt.Print("涂口红【")
	l.Decorator.Show()
	fmt.Print("】")
}
