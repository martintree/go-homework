package main

import (
	"fmt"
	"time"
)

func readDataFromChannel(ch <-chan int, done chan<- bool) {
	for data := range ch {
		fmt.Println("=====read data:", data)
	}
	done <- true
}

func sendDataToChannel(ch chan<- int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println(">>>>>send data:", i)
		if i%10 == 0 {
			time.Sleep(time.Second * 3)
		} else {
			time.Sleep(time.Millisecond * 100)
		}

	}
	close(ch) // 关闭通道
}
func main() {
	ch := make(chan int, 10)
	done := make(chan bool)
	go sendDataToChannel(ch)
	go readDataFromChannel(ch, done)

	for {
		if <-done {
			fmt.Println("get done signal")
			break
		}
	}
}
