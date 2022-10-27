package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch:
			fmt.Println("Message: ", m)
		case <-time.After(1 * time.Second):
			fmt.Println("1 sec timeout")
		}
	}
}
