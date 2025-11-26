package main

import (
	"fmt"
	"time"
)

type Task func()

func executeTask(name string, task Task, ch chan<- bool) {
	start := time.Now()
	task()
	elapsed := time.Since(start)
	fmt.Println("finished task:", name, "finished in:", elapsed)
	ch <- true
}

func main() {
	doneFlag1 := make(chan bool)
	doneFlag2 := make(chan bool)
	doneFlag3 := make(chan bool)

	go executeTask("task1", func() {
		time.Sleep(5 * time.Millisecond)
	}, doneFlag1)

	go executeTask("task2", func() {
		time.Sleep(2 * time.Millisecond)
	}, doneFlag2)

	go executeTask("task3", func() {
		time.Sleep(3 * time.Millisecond)
	}, doneFlag3)

	<-doneFlag1
	<-doneFlag2
	<-doneFlag3
}
