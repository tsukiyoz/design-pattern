package flyweight

import "fmt"

func ExampleFlyweight() {
	game := NewGame()

	// 添加恐怖分子玩家
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)

	// 添加反恐精英玩家
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)

	// 打印各个皮肤类型及颜色
	for dressType, dress := range GetDressFactory().store {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.Color())
	}
	// Output:
	// DressColorType: terrorist
	// DressColor: red
	// DressColorType: counterTerrorist
	// DressColor: green
}
