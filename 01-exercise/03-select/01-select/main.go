package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// TODO: multiplex recv on channel - ch1, ch2
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Message from channel 1: ", msg1)
		case msg2 := <-ch2:
			fmt.Println("Message from channel 2: ", msg2)
		}
	}

}
