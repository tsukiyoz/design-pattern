package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	startTime := time.Now()

	c := make(chan int)
	go fibonacci(15, c)

	for num := range c {
		fmt.Print(num, " ")
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("\nTotal time taken: %s\n", elapsedTime)
}
