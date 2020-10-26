package main

import "fmt"

// This func is used to return bool value. Disregard this func if you want to return majority element
func DivideAndConquer(arr []int32) bool {
	majorityElement := ifMajority(arr)
	if majorityElement == -1 {
		return false
	}
	return true
}

func ifMajority(arr []int32) int32 {
	n := int32(len(arr))
	if n == 1 {
		return arr[0]
	}
	leftArray := arr[:n/2]
	rightArray := arr[n/2 :]
	elemLeft := ifMajority(leftArray)
	elemRight := ifMajority(rightArray)

	// -1 value is referred to no-majority element in considered array.
	// Check -1 condition to reduce some calls to `getFrequency()` when return element is -1
	if elemLeft == -1 && elemRight == -1 {
		return -1
	}
	if elemRight == -1 || elemLeft == elemRight {
		return elemLeft
	}
	if elemLeft == -1 {
		return elemRight
	}
	freqLeft := getFrequency(arr, elemLeft)
	freqRight := getFrequency(arr, elemRight)
	if freqLeft > n/2 {
		return freqLeft
	} else if freqRight > n/2 {
		return freqRight
	} else {
		return -1
	}
}

// To get frequency of value k appears in arr
func getFrequency(arr []int32, k int32) int32 {
	count := int32(0)
	for _,v := range arr {
		if k == v {
			count += 1
		}
	}
	return count
}

func main() {
	//fmt.Println(ifMajority([]int32{2,2,3,3,4,2,5}))
	fmt.Println(DivideAndConquer([]int32{2,2,3,3,4,2,5}))
	//fmt.Println(ifMajority([]int32{2,3,3,3}))
	//fmt.Println(DivideAndConquer([]int32{2,3,3,3}))
}