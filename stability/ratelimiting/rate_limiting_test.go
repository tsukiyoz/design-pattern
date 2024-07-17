package ratelimit

import (
	"testing"
)

var (
	requestQueueSize = 10
)

func TestRateLimiting(t *testing.T) {
	//请求队列
	burstyRequests := make(chan int, requestQueueSize)

	for i := 1; i <= requestQueueSize; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	//对请求进行限流

	//200ms允许一次请求,最多同时3个突发
	rateLimiting(burstyRequests, 200, 3)
}
