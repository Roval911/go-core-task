package main

import (
	"fmt"
	"math"
)

func numberPipeline(in <-chan uint8, out chan<- float64) {
	for num := range in {
		out <- math.Pow(float64(num), 3)
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go numberPipeline(in, out)

	go func() {
		for i := uint8(1); i <= 10; i++ {
			in <- i
		}
		close(in)
	}()

	for result := range out {
		fmt.Printf("Result: %.2f\n", result)
	}
}
