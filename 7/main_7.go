package main

import (
	"fmt"
	"time"
)

func writeChan(start, stop int, ch chan int) {
	for i := start; i <= stop; i++ {
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func mergeChannels(channels ...chan int) chan int {
	mergedChannel := make(chan int)
	go func() {
		for _, ch := range channels {
			for val := range ch {
				mergedChannel <- val
			}
		}
		close(mergedChannel)
	}()
	return mergedChannel
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go writeChan(1, 5, ch1)
	go writeChan(6, 10, ch2)
	go writeChan(11, 10, ch3)
	merged := mergeChannels(ch1, ch2, ch3)
	for val := range merged {
		fmt.Println(val)
	}
}
