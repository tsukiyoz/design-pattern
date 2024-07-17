package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker", name, "is starting...")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker", name, "is done.")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go worker("A", &wg)
	go worker("B", &wg)
	go worker("C", &wg)

	wg.Wait()
	fmt.Println("All workers have finished.")
}
