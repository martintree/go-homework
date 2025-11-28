package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var counter int32 = 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("count:", counter)

}
