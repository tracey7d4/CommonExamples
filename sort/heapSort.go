package main

import "fmt"

type maxHeap struct {
	value []int32
}

func (m *maxHeap) leaf(index, size int32) bool {
	if index >= size/2 && index <= size {
		return true
	}
	return false
}

func (m *maxHeap) leftChild(index int32) int32 {
	return 2*index + 1
}

func (m *maxHeap) rightChild(index int32) int32 {
	return 2*index + 2
}

func (m *maxHeap) swap(i, j int32) {
	m.value[i], m.value[j] = m.value[j], m.value[i]
}

func (m *maxHeap) buildMaxHeap(size int32) {
	for i := size/2 - 1; i >= 0; i-- {
		m.siftDown(i, size)
	}
}

func (m *maxHeap) siftDown(index, size int32) {
	if m.leaf(index, size) {
		return
	}
	leftChildIndex := m.leftChild(index)
	rightChildIndex := m.rightChild(index)
	largest := index
	if leftChildIndex < size && m.value[leftChildIndex] > m.value[largest] {
		largest = leftChildIndex
	}
	if rightChildIndex < size && m.value[rightChildIndex] > m.value[largest] {
		largest = rightChildIndex
	}
	if largest != index {
		m.swap(index, largest)
		m.siftDown(largest, size)
	}
}

//func (m *maxHeap) heapSort(size int32) {
//	m.buildMaxHeap(size)
//	if size > 1 {
//		m.swap(0, size-1)
//		m.heapSort(size - 1)
//	}
//
//}
func (m *maxHeap) heapSort(size int32) {
	m.buildMaxHeap(size)
	for i := size - 1; i > 1; i-- {
		m.swap(0, size-1)
		m.siftDown(0, i)
	}

}

func main() {
	arr := []int32{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	m := new(maxHeap)
	for _, v := range arr {
		m.value = append(m.value, v)
	}
	m.heapSort(int32(len(arr)))
	fmt.Println(m.value)
}
