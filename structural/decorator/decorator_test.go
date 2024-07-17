package decorator

import "fmt"

func ExampleDecorator() {
	// 创建一个经过口红和粉底装饰的女生对象
	decoratedGirl := NewLipstick(NewFoundationMakeup(NewGirl()))
	decoratedGirl.Show() // 展示女生经过装饰后的样子
	fmt.Println()
	// Output:
	// 涂口红【打粉底【女生的素颜】】
}
