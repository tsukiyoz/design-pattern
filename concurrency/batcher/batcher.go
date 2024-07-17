package main

import (
	"fmt"
	"time"
)

const batchSize = 5

func processBatch(data []int) {
	fmt.Println("Processing batch:", data)
	// 模拟处理数据的逻辑
	time.Sleep(2 * time.Second)
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	totalData := len(data)

	for i := 0; i < totalData; i += batchSize {
		end := i + batchSize
		if end > totalData {
			end = totalData
		}
		batch := data[i:end]
		go processBatch(batch) // 并发处理批次数据
	}

	time.Sleep(3 * time.Second) // 等待异步处理完成
	fmt.Println("All batches processed.")
}
