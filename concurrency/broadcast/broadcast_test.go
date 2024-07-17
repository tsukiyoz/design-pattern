package broadcast

import (
	"fmt"
	"testing"
)

func TestBroadcast(t *testing.T) {
	receivers := make([]chan string, 3)

	// 创建三个接收者 Channel
	for i := range receivers {
		receivers[i] = make(chan string)
	}

	// 启动多个 Goroutines 作为接收者
	for i := range receivers {
		go func(id int, receiver chan string) {
			for {
				message := <-receiver // 接收消息
				fmt.Printf("Receiver %d got: %s\n", id, message)
			}
		}(i, receivers[i])
	}

	message := "Hello, this is a broadcast message"
	Broadcaster(message, receivers)

	// 等待一段时间
	select {}
}
