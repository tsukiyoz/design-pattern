package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("Task deadline:", deadline)
	}

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		fmt.Println("Task canceled due to deadline exceeded")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()

	fmt.Println("Task start")
	task(ctx)
	fmt.Println("Task end")
}
