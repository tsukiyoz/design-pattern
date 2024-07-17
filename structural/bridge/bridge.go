package bridge

import "fmt"

// Applications 家电接口定义了家用电器的基本行为
type Applications interface {
	Run()         // Run 方法表示家电运行的行为
	Name() string // Name 方法返回家电的名称
}

// Bulb 灯泡类实现了家电接口
type Bulb struct {
}

func (b *Bulb) Name() string {
	return "Bulb"
}

func (b *Bulb) Run() {
	fmt.Printf("灯泡 %v 亮了！\n", b.Name())
}

// Fan 电扇类实现了家电接口
type Fan struct {
}

func (f *Fan) Name() string {
	return "Fan"
}

func (f *Fan) Run() {
	fmt.Printf("电扇 %v 转动\n", f.Name())
}

// Switch 开关接口定义了开关的基本行为
type Switch interface {
	TurnOn()  // TurnOn 方法表示打开开关
	TurnOff() // TurnOff 方法表示关闭开关
}

// CircleSwitch 圆形开关类实现了开关接口
type CircleSwitch struct {
	application Applications // 关联的家电
}

func NewCircleSwitch(application Applications) *CircleSwitch {
	return &CircleSwitch{application: application}
}

func (c *CircleSwitch) TurnOn() {
	fmt.Printf("打开圆形开关，打开 %s，开始运行 ", c.application.Name())
	c.application.Run()
}

func (c *CircleSwitch) TurnOff() {
	fmt.Printf("关闭圆形开关，关闭 %s\n", c.application.Name())
}

// SquareSwitch 方形开关类实现了开关接口
type SquareSwitch struct {
	application Applications // 关联的家电
}

func NewSquareSwitch(application Applications) *SquareSwitch {
	return &SquareSwitch{application: application}
}

func (c *SquareSwitch) TurnOn() {
	fmt.Printf("打开方形开关，打开 %s，开始运行 ", c.application.Name())
	c.application.Run()
}

func (c *SquareSwitch) TurnOff() {
	fmt.Printf("关闭方形开关，关闭 %s\n", c.application.Name())
}
