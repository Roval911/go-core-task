package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateInt(ch chan int, rangeInt int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := rnd.Intn(rangeInt)
		ch <- num
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	go generateInt(ch, 100)
	for i := range ch {
		fmt.Println(i)
	}
}
