package main

import "fmt"

func main() {
	res := make(chan int)
	go func(a, b int) {
		c := a + b
		res <- c
	}(1, 2)
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", <-res)
}
