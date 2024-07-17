package main

import (
	"errors"
	"fmt"
	"os"
)

func processInput(input string) error {
	if input == "" {
		return errors.New("Input cannot be empty")
	}

	// 假设这里有一些处理逻辑

	return nil
}

func main() {
	input := ""
	err := processInput(input)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// 继续执行其他操作
	fmt.Println("Continue with other operations")
}
