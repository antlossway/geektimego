// use Channel to write a single-thread producer/consumer model
// channel cache 10, element type int
// producer: every second put a int in the channnel, chan will be blocked when it's full
// consumer: every second get a int from the chan, when chan empty, consumer is blocked

package main

import (
	"time"
	"math/rand"
)

func main() {

	ch := make(chan int, 10)
	go func() {
		for {
			j := <-ch
			println("get from chan:", j)
			time.Sleep(2 * time.Second)
		}
	}()


	for {
		i := rand.Intn(10000)
		ch <- i 
		println("put into chan: ", i)
		time.Sleep(time.Second)
	}	
}
