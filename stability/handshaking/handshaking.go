package main

import (
	"fmt"
	"time"
)

func clientHandshake(ch chan string) {
	fmt.Println("Client: Sending handshake request")
	ch <- "Handshake Request"
	response := <-ch
	fmt.Println("Client: Received handshake response -", response)
}

func serverHandshake(ch chan string) {
	fmt.Println("Server: Waiting for handshake request")
	request := <-ch
	fmt.Println("Server: Received handshake request -", request)
	time.Sleep(2 * time.Second) // 模拟处理时间
	fmt.Println("Server: Sending handshake response")
	ch <- "Handshake Response"
}

func main() {
	ch := make(chan string)

	go serverHandshake(ch)
	go clientHandshake(ch)

	time.Sleep(3 * time.Second) // 等待握手完成
}
