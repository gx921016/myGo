package main

import (
	"fmt"
	"time"
)

type worker struct {
	in   chan int
	done chan bool
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i)

	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].done
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
	time.Sleep(time.Millisecond)
}

func doWorker(id int, w worker) {

	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		w.done <- true
	}
}

//只能收数据	  <-chan int 只能发数据
func CreateWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w)

	return w
}

func main() {
	chanDemo()

}
