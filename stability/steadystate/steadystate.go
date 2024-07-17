package main

import (
	"fmt"
	"time"
)

func processTask(id int) {
	fmt.Printf("Processing task %d\n", id)
	time.Sleep(1 * time.Second) // 模拟任务处理时间
}

func main() {
	taskCount := 5
	taskInterval := 2 * time.Second

	fmt.Println("System is in Steady-State")

	for i := 1; i <= taskCount; i++ {
		go processTask(i)
		time.Sleep(taskInterval)
	}

	time.Sleep(taskInterval)
	fmt.Println("All tasks processed successfully")
}
