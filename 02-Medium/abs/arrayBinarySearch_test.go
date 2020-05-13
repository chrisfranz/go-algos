package main

import ( 
	"testing"
)

func TestArrayBinarySearch(t *testing.T) {
	got := arrayBinarySearch([]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 20, 0)
	want := -1

	if got != want {
		t.Error("try again")
	}
}