package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.
	// goroutine function call
	go fun("go routine direct call")

	// goroutine with anonymous function
	go func() {
		fun("go routine ananymous call")
	}()

	// goroutine with function value call
	go func(v string) {
		fun(v)
	}("go routine with value call")

	// wait for goroutines to end
	time.Sleep(1 * time.Second)

	fmt.Println("done..")
}
