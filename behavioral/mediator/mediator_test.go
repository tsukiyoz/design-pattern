package mediator

func ExampleMediator() {
	// 创建一个新的具体中介者实例
	mediator := NewConcreteMediator()

	// 创建三个用户参与者
	user1 := NewUser("Alice")
	user2 := NewUser("Bob")
	user3 := NewUser("Charlie")

	// 将用户参与者添加到中介者
	mediator.AddColleague(user1)
	mediator.AddColleague(user2)
	mediator.AddColleague(user3)

	// 为每个用户参与者设置中介者
	user1.SetMediator(mediator)
	user2.SetMediator(mediator)
	user3.SetMediator(mediator)

	// 发送消息给所有用户参与者
	user1.SendMessage("Hello, everyone!")
	user2.SendMessage("Hi, there!")
	user3.SendMessage("Hey, guys!")
	// Output:
	// [Bob] Received message: Hello, everyone!
	// [Charlie] Received message: Hello, everyone!
	// [Alice] Received message: Hi, there!
	// [Charlie] Received message: Hi, there!
	// [Alice] Received message: Hey, guys!
	// [Bob] Received message: Hey, guys!
}
