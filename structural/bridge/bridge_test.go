package bridge

func ExampleBridge() {
	// 使用圆形开关控制灯泡
	cs := NewCircleSwitch(&Bulb{})
	cs.TurnOn()
	cs.TurnOff()

	// 使用圆形开关控制电扇
	cs = NewCircleSwitch(&Fan{})
	cs.TurnOn()
	cs.TurnOff()

	// 使用方形开关控制灯泡
	ss := NewSquareSwitch(&Bulb{})
	ss.TurnOn()
	ss.TurnOff()

	// 使用方形开关控制电扇
	ss = NewSquareSwitch(&Fan{})
	ss.TurnOn()
	ss.TurnOff()
	// Output:
	// 打开圆形开关，打开 Bulb，开始运行 灯泡 Bulb 亮了！
	// 关闭圆形开关，关闭 Bulb
	// 打开圆形开关，打开 Fan，开始运行 电扇 Fan 转动
	// 关闭圆形开关，关闭 Fan
	// 打开方形开关，打开 Bulb，开始运行 灯泡 Bulb 亮了！
	// 关闭方形开关，关闭 Bulb
	// 打开方形开关，打开 Fan，开始运行 电扇 Fan 转动
	// 关闭方形开关，关闭 Fan
}
