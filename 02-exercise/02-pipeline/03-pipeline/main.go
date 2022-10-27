// generator() -> square() ->
//														-> merge -> print
//             -> square() ->
package main

import (
	"fmt"
	"sync"
)

func generator(done <-chan interface{}, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			select {
			case <-done:
				return
			case out <- n:
			}
		}
		close(out)
	}()
	return out
}

func square(done <-chan interface{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			select {
			case <-done:
				return
			case out <- n * n:
			}
		}
		close(out)
	}()
	return out
}

func merge(done <-chan interface{}, cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			select {
			case <-done:
				return
			case out <- n:
			}
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	done := make(chan interface{})
	defer close(done)

	in := generator(done, 2, 3, 55, 12, 58)

	c1 := square(done, in)
	c2 := square(done, in)

	out := merge(done, c1, c2)

	// TODO: cancel goroutines after receiving one value.

	fmt.Println(<-out)
}
