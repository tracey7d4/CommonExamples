package main

import (
	"fmt"
)

func quicksort2way(arr []int) []int {
	sort2way(arr, 0, len(arr)-1)
	return arr
}

func sort2way(arr []int, left, right int){
	pivot := arr[(right+left)/2]
	i := left
	j := right
	for i < j {
		for arr[i] < pivot && i <= j {
			i++
		}
		for arr[j] > pivot && i <= j {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	if left < i - 1 {
		sort2way(arr, left, i - 1)
	}
	if right > i {
		sort2way(arr, i, right)
	}
}

func main() {
	arr := []int{50,23,9,18,61,32}
	fmt.Println(quicksort2way(arr))
}