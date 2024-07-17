package main

import (
	"fmt"
	"time"
)

func serviceCall() error {
	// 模拟调用远程服务
	if time.Now().Second()%2 == 0 {
		return fmt.Errorf("Service Error")
	}
	return nil
}

func main() {
	breaker := false
	failCount := 0

	for i := 0; i < 10; i++ {
		if failCount >= 3 {
			breaker = true
			fmt.Println("Circuit-Breaker is open")
		}

		if breaker {
			fmt.Println("Service is unavailable")
			time.Sleep(2 * time.Second) // 等待断路器重置
			failCount = 0
			breaker = false
			continue
		}

		err := serviceCall()
		if err != nil {
			fmt.Println("Service call failed:", err)
			failCount++
		} else {
			fmt.Println("Service call successful")
		}
	}
}
