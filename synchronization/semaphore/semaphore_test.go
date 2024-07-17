package semaphore

import (
	"fmt"
	"testing"
	"time"
)

func TestSemaphoreAcquireRelease(t *testing.T) {
	sem := New(3, 1*time.Second)

	for i := 0; i < 10; i++ {
		if err := sem.Acquire(); err != nil {
			t.Error(err)
		}
		if err := sem.Acquire(); err != nil {
			t.Error(err)
		}
		if err := sem.Acquire(); err != nil {
			t.Error(err)
		}
		sem.Release()
		sem.Release()
		sem.Release()
	}
}

func TestSemaphoreBlockTimeout(t *testing.T) {
	sem := New(1, 200*time.Millisecond)

	if err := sem.Acquire(); err != nil {
		t.Error(err)
	}

	start := time.Now()
	if err := sem.Acquire(); err != ErrNoTickets {
		t.Error(err)
	}
	if start.Add(200 * time.Millisecond).After(time.Now()) {
		t.Error("semaphore did not wait long enough")
	}

	sem.Release()
	if err := sem.Acquire(); err != nil {
		t.Error(err)
	}
}

func TestSemaphoreEmpty(t *testing.T) {
	sem := New(2, 200*time.Millisecond)

	if !sem.IsEmpty() {
		t.Error("semaphore should be empty")
	}

	sem.Acquire()

	if sem.IsEmpty() {
		t.Error("semaphore should not be empty")
	}

	sem.Release()

	if !sem.IsEmpty() {
		t.Error("semaphore should be empty")
	}
}

func ExampleSemaphore() {
	sem := New(2, 1*time.Second)

	for i := 0; i < 5; i++ {
		go func() {
			if err := sem.Acquire(); err != nil {
				fmt.Println(err)
				return // 无法获取信号量
			}
			defer sem.Release()
			time.Sleep(2 * time.Second)

			return // 在信号量保护下执行某些操作
		}()
	}

	time.Sleep(5 * time.Second)
	// Output:
	// 无法获取信号量票证
	// 无法获取信号量票证
	// 无法获取信号量票证
}
