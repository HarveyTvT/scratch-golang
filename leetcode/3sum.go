package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	ret := make([][]int, 0)

	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]

		j := i + 1
		k := len(nums) - 1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			sum := nums[j] + nums[k]
			if sum == target {
				ret = append(ret, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
	}
	return ret
}
