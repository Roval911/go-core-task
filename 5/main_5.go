package main

func Intersection(slice1, slice2 []int) []int {
	var newSlice []int
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v1 == v2 {
				newSlice = append(newSlice, v1)
			}
		}
	}
	return newSlice
}
