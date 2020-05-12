package main

import "fmt"



func maxProfit(array []int) int {
	lowest := array[0]
	profit := 0

	for i := 1; i < len(array); i++ {
		if array[i] - lowest > profit {
			profit = array[i] - lowest
		}
		if array[i] < lowest {
			lowest = array[i]
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit([]int{10, 200, 32, 4025, 5,}))
}