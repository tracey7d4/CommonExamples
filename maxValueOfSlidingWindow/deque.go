package main

import (
	"fmt"
)

func main() {
	arr := []int32{7, 6, 5, 4, 8, 9, 1}
	fmt.Println(slidingWindow(arr, 3))
}

func slidingWindow(arr []int32, k int32) []int32 {
	n := int32(len(arr))
	dequeue := new(deque)
	dequeue.push(0)
	for i := int32(1); i < k; i++ {
		for len(dequeue.index) != 0 && arr[i] >= arr[dequeue.index[len(dequeue.index)-1]] {
			dequeue.popEnd()
		}
		dequeue.push(i)
	}
	res := make([]int32, 0, n)
	res = append(res, arr[dequeue.index[0]])
	for i := k; i < n; i++ {
		if dequeue.index[0] < i-k+1 {
			dequeue.popStart()
		}
		for len(dequeue.index) != 0 && arr[i] >= arr[dequeue.index[len(dequeue.index)-1]] {
			dequeue.popEnd()
		}
		dequeue.push(i)
		res = append(res, arr[dequeue.index[0]])
	}

	return res
}

type deque struct {
	index []int32
}

func (l *deque) popEnd() {
	n := len((*l).index)
	(*l).index = (*l).index[:n-1]
}

func (l *deque) popStart() {
	(*l).index = (*l).index[1:]
}

func (l *deque) push(a int32) {
	(*l).index = append((*l).index, a)
}

