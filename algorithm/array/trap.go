package array

//https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	ans, leftMax, rightMax := 0, 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += (leftMax - height[left])
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += (rightMax - height[right])
			}
			right--
		}
	}
	return ans
}
