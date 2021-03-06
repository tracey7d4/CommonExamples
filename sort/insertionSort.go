package main

import "fmt"

//
func insertionSort(arr []int32) []int32 {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}

func main() {
	arr := []int32{1, 6, 4, 2, 7, 8, 5}
	fmt.Println(insertionSort(arr))
}
