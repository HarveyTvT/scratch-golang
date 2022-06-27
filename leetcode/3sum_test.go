// https://leetcode.cn/problems/3sum/
package leetcode

import (
	"reflect"
	"testing"
)

func Test3Sum(t *testing.T) {
	input := [][]int{
		{-1, 0, 1, 2, -1, -4},
		{},
		{0},
	}

	expected := [][][]int{
		{[]int{-1, -1, 2}, []int{-1, 0, 1}},
		{},
		{},
	}

	for i := range input {
		got := ThreeSum(input[i])
		want := expected[i]

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

}
