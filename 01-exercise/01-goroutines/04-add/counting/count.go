package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	// Utilize all cores on machine
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Divide the input into parts
	var wg sync.WaitGroup
	ch := make(chan int64)

	length := len(numbers)
	j := runtime.NumCPU() // divide into n parts, where n is number of cores

	// Run computation for each part in seperate goroutine.

	go func() {

		for i := 0; i < j; i++ {
			wg.Add(1)
			part := numbers[i*length/j : (i+1)*length/j]
			go func(nums []int) {
				defer wg.Done()
				sum := Add(nums)
				ch <- sum
			}(part)

		}

		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		sum += v
	}
	// Add part sum to cummulative sum

	return sum
}
