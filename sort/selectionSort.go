package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIndex := findIndexMin(arr, i)
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

func findIndexMin(arr []int, start int) int {
	i := start
	min := arr[start]
	j := start
	for j < len(arr)-1 {
		j++
		if min > arr[j] {
			min = arr[j]
			i = j
		}
	}
	return i
}

func main() {
	arr := []int{6, 4, 2, 7, 5}
	fmt.Println(selectionSort(arr))
}