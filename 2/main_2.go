package main

import (
	"fmt"
	"math/rand"
	"time"
)

func OriginalSlice(rangeInt, len int) []int {
	originalSlice := make([]int, len)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range originalSlice {
		originalSlice[i] = rnd.Intn(rangeInt)
	}
	return originalSlice
}

func sliceExample(in []int) []int {
	var out []int
	for i := range in {
		if in[i]%2 == 0 {
			out = append(out, in[i])
		}
	}
	return out
}

func addElements(in []int, n int) []int {
	return append(in, n)
}

func copySlice(orig []int) []int {
	duplicate := make([]int, len(orig))
	copy(duplicate, orig)
	return duplicate
}

func removeElement(orig []int, index int) []int {
	return append(orig[:index], orig[index+1:]...)
}

func main() {
	original := OriginalSlice(100, 10)
	fmt.Println("Оригинальный слайс:", original)

	copied := copySlice(original)
	fmt.Println("Скопированный слайс:", copied)

	filtered := sliceExample(original)
	fmt.Println("Слайс с четными числами:", filtered)

	added := addElements(original, 42)
	fmt.Println("Слайс после добавления элемента:", added)

	if len(original) > 2 {
		removed := removeElement(original, 2)
		fmt.Println("Слайс после удаления элемента по индексу 2:", removed)
	} else {
		fmt.Println("Недостаточно элементов для удаления по индексу 2")
	}
}
