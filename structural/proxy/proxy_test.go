package proxy

func ExampleProxy() {
	foursMarket := NewFoursMarketProxy(NewUser("小明"))
	foursMarket.BuyCar()
	// Output:
	// 汽车上牌子
	// 注册汽车
	// 从供应商购买汽车到 4S 店
	// <小明>买车
	// 保养汽车
	// 维修汽车
}
