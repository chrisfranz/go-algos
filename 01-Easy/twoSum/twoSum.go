package main

import (
	"fmt"
	"sort"
)

func twoSum(array []int, target int) []int {
	nums := map[int]bool{}

	for i := 0; i < len(array); i++ {
		if !nums[array[i]] {
			nums[target - array[i]] = true;
		} else {
			return []int{ array[i], target - array[i] }
		}
	}
	return []int{};
}

func twoSumSorted(array []int, target int) []int {
	sort.Ints(array)

	left := 0
	right := len(array) - 1
	
	for left < right {
		leftVal := array[left]
		rightVal := array[right]
		sum := leftVal + rightVal

		if sum == target {
			return []int{ leftVal, rightVal}
		} else if sum < target { 
			left++
		} else {
			right--
		}
	}
	return[]int{}
}

// var output1  = twoSum([]int{7, 6, 5, 4}, 13)


func main() {
	var output2  = twoSumSorted([]int{7, 6, 5, 4, -20, 234, 9, 22, -7}, -1690)
	fmt.Println(output2)
}
