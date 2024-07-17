package pubsub

import (
	"fmt"
)

// Publisher 结构体表示发布者，包含订阅者的映射
type Publisher struct {
	subscribers map[string]chan string
}

// NewPublisher 创建一个新的发布者实例
func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[string]chan string),
	}
}

// Subscribe 订阅指定主题的消息
func (p *Publisher) Subscribe(topic string, ch chan string) {
	p.subscribers[topic] = ch
}

// Unsubscribe 取消订阅指定主题的消息
func (p *Publisher) Unsubscribe(topic string) {
	delete(p.subscribers, topic)
}

// Publish 发布消息到指定主题
func (p *Publisher) Publish(topic, message string) {
	if ch, ok := p.subscribers[topic]; ok {
		ch <- message
	}
}

func main() {
	publisher := NewPublisher()

	ch1 := make(chan string)
	ch2 := make(chan string)

	publisher.Subscribe("topic1", ch1)
	publisher.Subscribe("topic2", ch2)

	// 启动订阅者的 goroutine 接收消息
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

	// 发布消息到指定主题
	publisher.Publish("topic1", "Hello, World!")
	publisher.Publish("topic2", "Greetings from Publisher!")
}
