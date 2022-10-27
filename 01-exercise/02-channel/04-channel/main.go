package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch1 chan string, msgs ...string) {
	// send message on ch1
	for _, msg := range msgs {
		ch1 <- msg
	}
	close(ch1)
}

func relayMsg(ch1, ch2 chan string) {
	// recv message on ch1
	// send it on ch2

	for v := range ch1 {
		ch2 <- v
	}

	close(ch2)
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spine goroutine genMsg and relayMsg
	go genMsg(ch1, "first msg", "second msg", "third msg")
	go relayMsg(ch1, ch2)

	// recv message on ch2

	for msg := range ch2 {
		fmt.Println(msg)
	}
}
