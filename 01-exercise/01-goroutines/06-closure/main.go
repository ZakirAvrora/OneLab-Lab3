package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
		}(i)
	}
	wg.Wait()
}
