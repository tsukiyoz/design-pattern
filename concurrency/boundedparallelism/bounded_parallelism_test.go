package boundedparallelism

import (
	"fmt"
	"sync"
	"testing"
)

func TestBoundedParallelism(t *testing.T) {
	numWorkers := 5
	boundedParallelism := 2
	semaphore := make(chan struct{}, boundedParallelism)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(id int) {
			defer wg.Done()
			Worker(id, semaphore)
		}(i)
	}

	wg.Wait()

	fmt.Println("All workers have finished, continue to the next step")
}
