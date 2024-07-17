package proxy

import "fmt"

// 买车行为接口
type IBuyCar interface {
	BuyCar()
}

// 用户买车
type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func (u *User) BuyCar() {
	fmt.Printf("<%s>买车\n", u.Name)
}

// 4S 店代理帮你买车
type FoursMarketProxy struct {
	User *User
}

func NewFoursMarketProxy(user *User) *FoursMarketProxy {
	return &FoursMarketProxy{User: user}
}

func (f *FoursMarketProxy) BuyCar() {
	fmt.Println("汽车上牌子")
	fmt.Println("注册汽车")
	fmt.Println("从供应商购买汽车到 4S 店")
	f.User.BuyCar()
	fmt.Println("保养汽车")
	fmt.Println("维修汽车")
}
