package main

import ( 
	"testing"
	"reflect"
)

func TestTwoSum(t *testing.T) {
	got := twoSum([]int{ 7, 6, 5, 4}, 13)
	want := []int{ 7, 6 }

	if reflect.DeepEqual(got, want) {
		t.Error("try again")
	}
}