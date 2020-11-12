// Find the subarray from given array of integers so that sum of all elements in the subarray is maximum
package main

import "fmt"

// arr - given array
// return: sum of all elements in the subarray && subarray
func maxSubarray(arr []int32) (int32, []int32) {
	// In case elements of Array are all negative, return maximum element
	maxArray := arr[0]
	index := 0
	for i, v := range arr {
		if v > maxArray {
			maxArray = v
			index = i
		}
	}
	if maxArray < 0 {
		return maxArray, arr[index : index+1]
	}

	maxSum := int32(0)
	currentSum := int32(0)
	i := 0 //start index of largest subarray
	j := 1 // end index of largest subarray
	tempIndex := 0
	for k, v := range arr {
		currentSum += v
		if currentSum >= 0 {
			if maxSum <= currentSum {
				maxSum = currentSum
				i = tempIndex
				j = k + 1
			}
		} else {
			tempIndex = k + 1
			currentSum = 0
		}
	}

	return maxSum, arr[i:j]
}

func main() {
	//arr := []int32{-1, 0, 0, -2}
	arr := []int32{2, -1, -2, 3, 4, -5}
	//arr := []int32{2, -1, 2, 3, 4, -5}
	maxSum, subarr := maxSubarray(arr)
	fmt.Println(maxSum)
	fmt.Println(subarr)
}
