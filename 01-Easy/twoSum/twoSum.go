package main

import "fmt"

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

func main() {
	fmt.Println(twoSum([]int{7,6,5,4}, 13))
}
