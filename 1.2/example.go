package main

import (
	"fmt"
	"time"
)

func producer(ch chan <-int) {
	for i:= 0; i < 10; i++ {
		ch <- i
		fmt.Printf("produce data: %d\n", i)
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func consumer(ch <- chan int) {
	for i :=0; i < 10; i++ {
		data := <- ch
		fmt.Printf("fetch data: %d\n", data)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 5)
	go  producer(ch)
	consumer(ch)
}
