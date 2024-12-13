package main

func uniqSlice(slice1, slice2 []string) []string {
	var newSlice []string
	for _, v1 := range slice1 {
		flag := false
		for _, v2 := range slice2 {
			if v1 == v2 {
				flag = true
				break
			}
		}
		if !flag {
			newSlice = append(newSlice, v1)
		}
	}
	return newSlice
}
