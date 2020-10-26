package main

import "fmt"

func useMap(arr []int32) bool {
	n := int32(len(arr))
	m := make(map[int32]int32)
	for _,v := range arr {
		m[v] += 1
		if m[v] > n/2 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(useMap([]int32{2,2,2,3,4,2}))
	fmt.Println(useMap([]int32{2,2,3}))
	fmt.Println(useMap([]int32{1}))
}