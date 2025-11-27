package main

import (
	"fmt"
	"sync"
)

// 定义结构体
type Counter struct {
	Count int
	Mtx   sync.Mutex
}

func (c *Counter) Increment() {
	c.Mtx.Lock()
	defer c.Mtx.Unlock()
	c.Count++
}

func (c *Counter) GetCount() int {
	c.Mtx.Lock()
	defer c.Mtx.Unlock()
	return c.Count
}
func main() {
	counter := &Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter.GetCount())
}
