// +build ignore

package main

import (
	"fmt"
	"github.com/m3hm3t/go-concurrency-patterns-library/pkg/split"
	"sync"
)

func main() {
	source := make(chan int)                // The input channel
	dests := split.Split(source, 5)               // Retrieve 5 output channels

	go func() {                             // Send the number 1..10 to source
		for i := 1; i <= 10; i++ {          // and close it when we're done
			source <- i
		}

		close(source)
	}()

	var wg sync.WaitGroup                   // Use WaitGroup to wait until
	wg.Add(len(dests))                      // the output channels all close

	for i, ch := range dests {
		go func(i int, d <-chan int) {
			defer wg.Done()

			for val := range d {
				fmt.Printf("#%d got %d\n", i, val)
			}
		}(i, ch)
	}

	wg.Wait()
}

