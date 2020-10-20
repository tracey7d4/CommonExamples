package main

import "fmt"

// bubbleSort O(n^2)
func bubbleSort(arr []int32) []int32 {
	n := len(arr)
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
	return arr
}

func main() {
	arr := []int32{6, 4, 2, 7, 5}
	fmt.Println(bubbleSort(arr))
}
