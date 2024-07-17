package main

import (
	"fmt"
	"time"
)

func retryOperation(operation func() error, maxRetries int, interval time.Duration) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = operation()
		if err == nil {
			return nil // 操作成功，直接返回
		}
		fmt.Printf("Operation failed: %s. Retrying...\n", err)
		time.Sleep(interval)
	}
	return err
}

func main() {
	maxRetries := 3
	interval := 2 * time.Second

	err := retryOperation(func() error {
		fmt.Println("Performing operation...")
		return performOperation() // 假设 performOperation 是一个可能失败的操作
	}, maxRetries, interval)

	if err != nil {
		fmt.Printf("Operation failed after %d retries: %s\n", maxRetries, err)
	} else {
		fmt.Println("Operation successful!")
	}
}

func performOperation() error {
	// 模拟一个可能失败的操作
	if time.Now().Unix()%2 == 0 {
		return fmt.Errorf("Random error occurred")
	}
	return nil
}
