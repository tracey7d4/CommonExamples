package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 1, 2, 2, 3, 3}

	fmt.Println(partitionArray(arr))
}

func partitionArray(arr []int) bool {
	n := len(arr)
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if n < 3 || sum%3 != 0 {
		return false
	}
	k := sum / 3

	m := make(map[string]bool)
	return threeSubset(arr, k, k, k, m)
}

func threeSubset(arr []int, k1, k2, k3 int, m map[string]bool) bool {
	n := len(arr)
	if k1 == 0 && k2 == 0 && k3 == 0 {
		return true
	}
	if n == 0 {
		return false
	}
	s := strconv.Itoa(k1) + "+" + strconv.Itoa(k2) + "+" + strconv.Itoa(k3) + "+" + strconv.Itoa(n)
	if _,ok := m[s]; ok {
		return m[s]
	}
	a := false
	if k1 >= arr[0] {
		a = threeSubset(arr[1:], k1-arr[0], k2, k3, m)
	}
	if a == true {
		m[s] = true
		return a
	}
	b := false
	if k2 >= arr[0] {
		b = threeSubset(arr[1:], k1, k2-arr[0], k3, m)
	}
	if b == true {
		m[s] = true
		return b
	}
	c := false
	if k2 >= arr[0] {
		c = threeSubset(arr[1:], k1, k2, k3-arr[0], m)
	}
	m[s] = c
	return c

}
