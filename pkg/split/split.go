package split

func Split(source <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0)          // Create the dests slice

	for i := 0; i < n; i++ {                // Create n destination channels
		ch := make(chan int)
		dests = append(dests, ch)

		go func() {                         // Each channel gets a dedicated
			defer close(ch)                 // goroutine that competes for reads

			for val := range source {
				ch <- val
			}
		}()
	}

	return dests
}