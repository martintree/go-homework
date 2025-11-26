package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println("odd number:", i)
		}
	}()

	go func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println("even number:", i)
		}
	}()

	time.Sleep(10 * time.Millisecond)
}
