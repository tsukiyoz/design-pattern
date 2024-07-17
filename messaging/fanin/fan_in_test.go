package main

import (
	"fmt"
)

func ExampleFanIn() {
	// 第一路输入源
	dataStreams1 := []int{13, 44, 56, 99, 9, 45, 67, 90, 78, 23}
	// 生成带有输入的通道
	inputChan1 := generateNumbersPipeline(dataStreams1)

	// 第二路输入源
	dataStreams2 := []int{2, 4, 6, 9, 1, 1, 2, 3, 7, 8}
	inputChan2 := generateNumbersPipeline(dataStreams2)

	// 对每个输入通道进行平方操作
	c1 := squareNumber(inputChan1)
	c2 := squareNumber(inputChan2)

	// 将两路数据的平方数进行合并
	out := Merge(c1, c2)

	sum := 0

	// 计算合并后的平方数的总和
	for c := range out {
		sum += c
	}

	fmt.Printf("FanIn 合并后的平方数总和为: %d\n", sum)
	// Output:
	// FanIn 合并后的平方数总和为: 36615
}
