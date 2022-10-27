package main

import (
	"fmt"
	"math/rand"
	"time"
)

//TODO: identify the data race
// fix the issue.

func main() {
	start := time.Now()
	var t *time.Timer
	reset := make(chan struct{})

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		reset <- struct{}{}
	})

	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------
