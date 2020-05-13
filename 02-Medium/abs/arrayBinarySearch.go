package main

import (
	"fmt"
)

func arrayBinarySearch(array []int, target int, startVal int) int {
	midpoint := len(array) / 2
	
	if len(array) == 0 {
		return -1
	}

	if array[midpoint] == target {
		return startVal + midpoint
	}

	if array[midpoint] > target {
		return arrayBinarySearch(array[0:midpoint], target, startVal)
	}

	if array[midpoint] < target {
		startVal = startVal + midpoint + 1
		return arrayBinarySearch(array[midpoint + 1:], target, startVal)
	}
	return -1
}

func main() {
	fmt.Println(arrayBinarySearch([]int{1,2,3,4,5,6,7,8,9}, 7, 0))
}