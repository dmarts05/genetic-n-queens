package util

import (
	"slices"
	"testing"
)

func TestSample(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 5
	got := Sample(numbers, n)
	// Check if the length of the result is the same as the input
	if len(got) != n {
		t.Errorf("Sample() = %v, want %v", len(got), n)
	}
	// Check if the result contains only elements from the input
	for _, v := range got {
		if !slices.Contains(numbers, v) {
			t.Errorf("Sample() = %v, want %v", got, numbers)
		}
	}
	// Check whether the result contains only unique elements
	for i, v := range got {
		for j, v2 := range got {
			if i != j && v == v2 {
				t.Errorf("Sample() = %v, want unique elements", got)
			}
		}
	}
}
