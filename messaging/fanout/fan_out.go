package fanout

import "sync"

// Split 将一个通道分发到多个通道，数据轮流分发
func Split(ch <-chan int, n int) []chan int {
	cs := []chan int{}
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	distributeToChannels := func(ch <-chan int, cs []chan int) {
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()

		for {
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}
				for _, c := range cs {
					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}

// Split2 将一个通道分发到多个通道，每个 worker 分发一轮数据
func Split2(ch <-chan int, n int) []chan int {
	cs := []chan int{}
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	var wg sync.WaitGroup
	distributeToChannels := func(ch <-chan int, cs []chan int) {
		val, ok := <-ch
		if !ok {
			return
		}
		for _, c := range cs {
			c <- val
		}
		wg.Done()
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go distributeToChannels(ch, cs)
	}

	go func() {
		wg.Wait()
		for _, c := range cs {
			close(c)
		}
	}()

	return cs
}

// Split3 将一个通道分发到多个通道，随机分发数据到不同通道
func Split3(ch <-chan int, n int) []chan int {
	cs := make([]chan int, n)
	for i := 0; i < n; i++ {
		cs[i] = make(chan int)
	}

	distributeToChannels := func(ch <-chan int, cs []chan int) {
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}
					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}

// gen 函���将一组整数转换为一个通道，按顺序发送整数并在发送完毕后关闭通道
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// sq 函数接收一个整数通道，返回一个按平方处理过的整数通道，并在处理完毕后关闭通道
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
