/**
 * created by
 * @project advanced-star
 * @author frankstar
 * @date 2023/2/12
 * @contact frankstarye@tencent.com
 **/

package medium

func maxArea(height []int) int {
	l, r := 0, len(height) - 1
	square := 0
	for l < r {
		tmp := (r - l) * minInt(height[l], height[r])
		square = maxInt(square, tmp)
		if height[l] > height[r] {
			r --
		} else {
			l ++
		}
	}
	return square
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

