package flyweight

import "fmt"

// Dress 享元接口，定义了获取皮肤颜色的方法
type Dress interface {
	Color() string
}

// TerroristDress 具体的恐怖分子皮肤对象，实现了 Dress 接口
type TerroristDress struct {
	color string
}

// NewTerroristDress 实例化恐怖分子皮肤对象
func NewTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"} // 红色皮肤
}

// Color 获取恐怖分子皮肤的颜色
func (t *TerroristDress) Color() string {
	return t.color
}

// CounterTerroristDress 具体的反恐精英皮肤对象，实现了 Dress 接口
type CounterTerroristDress struct {
	color string
}

// NewCounterTerroristDress 实例化反恐精英皮肤对象
func NewCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"} // 绿色皮肤
}

// Color 获取反恐精英皮肤的颜色
func (t *CounterTerroristDress) Color() string {
	return t.color
}

// 享元的工厂
const (
	// 恐怖分子皮肤类型
	TerroristDressType = "terrorist"
	// 反恐精英皮肤类型
	CounterTerroristDressType = "counterTerrorist"
)

// DressFactory 皮肤工厂，负责管理不同类型的皮肤对象
type DressFactory struct {
	store map[string]Dress
}

// GetDress 根据皮肤类型获取皮肤对象
func (d *DressFactory) GetDress(dress string) (Dress, error) {
	// 如果已经创建过对象就直接返回，实现享元模式，减少对象创建
	if d.store[dress] != nil {
		return d.store[dress], nil
	}

	// 根据皮肤类型实例化相应皮肤对象
	if dress == TerroristDressType {
		d.store[dress] = NewTerroristDress()
		return d.store[dress], nil
	}

	if dress == CounterTerroristDressType {
		d.store[dress] = NewCounterTerroristDress()
		return d.store[dress], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

var dressFactory = &DressFactory{store: make(map[string]Dress)}

// GetDressFactory 获取皮肤工厂实例
func GetDressFactory() *DressFactory {
	return dressFactory
}

// Player 玩家结构体，包含玩家类型和所使用的皮肤
type Player struct {
	dress      Dress
	playerType string // 玩家类型，T代表恐怖分子，CT代表反恐精英
}

// newPlayer 创建新的玩家实例
func newPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactory().GetDress(dressType)
	return &Player{playerType: playerType, dress: dress}
}

// Game 游戏结构体，包含恐怖分子和反恐精英玩家列表
type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

// NewGame 创建新的游戏实例
func NewGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 0),
		counterTerrorists: make([]*Player, 0),
	}
}

// AddTerrorist 向游戏中添加恐怖分子玩家
func (g *Game) AddTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
	return
}

// AddCounterTerrorist 向游戏中添加反恐精英玩家
func (g *Game) AddCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
	return
}
