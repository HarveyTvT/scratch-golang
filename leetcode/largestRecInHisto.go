// https://leetcode.cn/problems/largest-rectangle-in-histogram/
package leetcode

func LargestRectangleArea(heights []int) int {
	f := make([][]int, len(heights))
	for i := range f {
		f[i] = make([]int, len(heights))
		f[i][i] = heights[i]
	}

	i := 0
	j := len(heights) - 1

	ret := 0

	for i <= j {
		step := j - i + 1
		height := get(heights, f, i, j)
		area := step * height
		if area > ret {
			ret = area
		}

		if heights[j] > heights[i] {
			i++
		} else {
			j--
		}
	}

	return ret

}

func get(heights []int, f [][]int, i, j int) int {
	if f[i][j] > 0 {
		return f[i][j]
	}

	if i == j {
		f[i][j] = heights[i]
		return f[i][j]
	}

	f[i][j] = min(heights[j], get(heights, f, i, j-1))
	return f[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
