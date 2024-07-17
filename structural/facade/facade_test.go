package facade

func ExampleFacade() {
	facade := NewFacade()
	facade.Order()
	// Output:
	// 开始为用户服务
	// 下厨烹饪
	// 开始为用户服务
	// 开始清洁卫生
}
