package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		for _, e := range nums {
			intStream <- e
		}
	}()

	return intStream
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()

	return out
}

func main() {
	// set up the pipeline
	in := generator(1, 2, 3, 4, 5, 6)

	// run the last stage of pipeline
	out := square(in)

	// receive the values from square stage
	// print each one, until channel is closed.
	for v := range out {
		fmt.Println(v)
	}

}
