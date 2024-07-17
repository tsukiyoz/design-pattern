package pubsub

import (
	"fmt"
	"testing"
)

func TestPublishSubscribe(t *testing.T) {
	// 创建一个新的发布者实例
	publisher := NewPublisher()

	// 创建两个消息通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 订阅两个不同的主题，并将消息通道传入
	publisher.Subscribe("topic1", ch1)
	publisher.Subscribe("topic2", ch2)

	// 启动一个 goroutine 来接收订阅的消息
	go func() {
		for {
			select {
			case msg := <-ch1:
				fmt.Println("Subscriber 1 received message:", msg)
			case msg := <-ch2:
				fmt.Println("Subscriber 2 received message:", msg)
			}
		}
	}()

	// 发布两条消息到不同的主题
	publisher.Publish("topic1", "Hello, World!")
	publisher.Publish("topic2", "Greetings from Publisher!")
}
