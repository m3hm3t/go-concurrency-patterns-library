// +build ignore

package main

import (
	"fmt"
	"github.com/m3hm3t/go-concurrency-patterns-library/pkg/funnel"
	"time"
)

func main() {

	sources := make([]<-chan int, 0)        // Create an empty channel slice

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch)       // Create a channel; add to sources

		go func() {                         // Run a toy goroutine for each
			defer close(ch)                 // Close ch when the routine ends

			for i := 1; i <= 5; i++ {
				ch <- i
				time.Sleep(time.Second)
			}
		}()
	}

	dest := funnel.Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}
}

