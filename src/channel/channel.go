package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)

	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

//只能收数据	  <-chan int 只能发数据
func CreateWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)

	return c
}

func worker(id int, c chan int) {
	//for true {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n",
	//		id, n)
	//}
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}
func main() {
	//chanDemo()
	//bufferedChannel()
	//channelClose()
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	close(c)
	time.Sleep(time.Millisecond)
}
