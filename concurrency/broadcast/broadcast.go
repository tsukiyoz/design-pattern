package broadcast

func Broadcaster(message string, receivers []chan string) {
	for _, receiver := range receivers {
		receiver <- message // 将消息发送给每个接收者
	}
}
