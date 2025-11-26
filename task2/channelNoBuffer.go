package main

import (
	"fmt"
	// "time"
)

func readData(ch <-chan int, done chan<- bool) {
	for data := range ch {
		fmt.Println("read data:", data)
	}
	done <- true
}

func sendData(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // 关闭通道
}
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sendData(ch)
	go readData(ch,done)

	 for {
		if <-done {
			fmt.Println("get done signal")
			break
		}
	}
}
