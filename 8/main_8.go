package main

import (
	"fmt"
	"sync"
	"time"
)

type SimpleWaitGroup struct {
	counter int
	mu      sync.Mutex
	ch      chan struct{}
}

func (swg *SimpleWaitGroup) Add(delta int) {
	swg.mu.Lock()
	defer swg.mu.Unlock()
	if swg.counter == 0 && delta > 0 {
		swg.ch = make(chan struct{})
	}
	swg.counter += delta
}

func (swg *SimpleWaitGroup) Done() {
	swg.mu.Lock()
	defer swg.mu.Unlock()
	swg.counter--
	if swg.counter == 0 {
		close(swg.ch)
	}
}

func (swg *SimpleWaitGroup) Wait() {
	<-swg.ch
}

func main() {
	var swg SimpleWaitGroup
	const numGoroutines = 5

	for i := 1; i <= numGoroutines; i++ {
		swg.Add(1)
		go func(id int) {
			defer swg.Done()
			time.Sleep(time.Millisecond * time.Duration(100*id))
			fmt.Printf("Горутина %d выполнена\n", id)
		}(i)
	}

	swg.Wait()
	fmt.Println("ВСе Горутины закончили выполнение")
}
