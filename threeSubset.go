package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 1, 2, 2, 3, 3}

	fmt.Println(partitionSouvenirs(arr))
}

func partitionSouvenirs(arr []int) int {
	n := len(arr)
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if n < 3 || sum%3 != 0 {
		return 0
	}
	k := sum / 3

	m := make(map[string]int)
	return threeSubset(arr, k, k, k, m)
}

func threeSubset(arr []int, k1, k2, k3 int, m map[string]int) int {
	n := len(arr)
	if k1 == 0 && k2 == 0 && k3 == 0 {
		return 1
	}
	if n == 0 {
		return 0
	}
	s := strconv.Itoa(k1) + "+" + strconv.Itoa(k2) + "+" + strconv.Itoa(k3) + "+" + strconv.Itoa(n)
	if _,ok := m[s]; ok {
		return m[s]
	}
	a := 0
	if k1 >= arr[0] {
		a = threeSubset(arr[1:], k1-arr[0], k2, k3, m)
	}
	if a == 1 {
		m[s] = 1
		return a
	}
	b := 0
	if k2 >= arr[0] {
		b = threeSubset(arr[1:], k1, k2-arr[0], k3, m)
	}
	if b == 1 {
		m[s] = 1
		return b
	}
	c := 0
	if k2 >= arr[0] {
		c = threeSubset(arr[1:], k1, k2, k3-arr[0], m)
	}
	m[s] = c
	return c

}
