package main

import (
	"fmt"
)

func main() {
	arr := []int32{7, 6, 5, 4, 8, 9, 1}
	fmt.Println(slidingWindow(arr, 3))
}

func slidingWindow(arr []int32, m int32) []int32 {
	n := int32(len(arr))
	maxIndex := new(list)
	maxIndex.push(0)
	for i := int32(1); i < m; i++ {
		for len(maxIndex.index) != 0 && arr[i] >= arr[maxIndex.index[len(maxIndex.index)-1]] {
			maxIndex.popEnd()
		}
		maxIndex.push(i)
	}
	res := make([]int32, 0, n)
	res = append(res, arr[maxIndex.index[0]])
	si := int32(1)
	for i := m; i < n; i++ {
		if maxIndex.index[0] < si {
			maxIndex.popStart()
		}
		for len(maxIndex.index) != 0 && arr[i] >= arr[maxIndex.index[len(maxIndex.index)-1]] {
			maxIndex.popEnd()
		}
		maxIndex.push(i)
		res = append(res, arr[maxIndex.index[0]])
		si++
	}

	return res
}

type list struct {
	index []int32
}

func (l *list) popEnd() {
	n := len((*l).index)
	(*l).index = (*l).index[:n-1]
}

func (l *list) popStart() {
	(*l).index = (*l).index[1:]
}

func (l *list) push(a int32) {
	(*l).index = append((*l).index, a)
}

