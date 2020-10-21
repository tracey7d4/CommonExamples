// https://www.hackerrank.com/challenges/pairs/problem

package main

import "fmt"

func pairs(k int32, arr []int32) int32 {
	m := make(map[int32]int32)
	for _, v := range arr {
		m[v] += 1
	}

	pairs := int32(0)
	for _, v := range arr {
		if _, ok := m[v+k]; ok {
			pairs += 1
		}
	}
	return pairs
}

func main() {
	arr := []int32{1, 5, 3, 4, 2}
	fmt.Println(pairs(2, arr))
}
