package main

import (
	"fmt"
	"time"
)

func componentA() error {
	// 模拟组件A的任务
	time.Sleep(1 * time.Second)
	return fmt.Errorf("Error in component A")
}

func componentB() error {
	// 模拟组件B的任务
	time.Sleep(2 * time.Second)
	return fmt.Errorf("Error in component B")
}

func main() {
	errA := componentA()
	if errA != nil {
		fmt.Println("Component A failed:", errA)

		// 处理组件A的故障，避免影响到组件B
		errB := componentB()
		if errB != nil {
			fmt.Println("Component B failed:", errB)
		}
	}
}
