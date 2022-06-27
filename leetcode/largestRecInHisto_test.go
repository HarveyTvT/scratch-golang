package leetcode

import (
	"testing"
)

func Test(t *testing.T) {
	input := [][]int{
		{2, 1, 5, 6, 2, 3},
		{0, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	expected := []int{
		10,
		20,
	}

	for i := range input {
		got := LargestRectangleArea(input[i])
		want := expected[i]

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
