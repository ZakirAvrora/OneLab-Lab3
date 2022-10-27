// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	var wg sync.WaitGroup
	multiplexStream := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for v := range c {
			multiplexStream <- v
		}
	}

	for _, c := range cs {
		wg.Add(1)
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexStream)
	}()

	return multiplexStream
}

func main() {
	in := generator(2, 3, 10, 12, 8, 5)

	// TODO: fan out square stage to run two instances.
	numSquares := 2
	squares := make([]<-chan int, numSquares)
	for i := 0; i < numSquares; i++ {
		squares[i] = square(in)
	}

	// TODO: fan in the results of square stages.
	res := merge(squares...)

	for v := range res {
		fmt.Println(v)
	}
}
