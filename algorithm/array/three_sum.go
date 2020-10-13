package array

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)

	numMap := make(map[int]int)
	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		numMap[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := -nums[i] - nums[j]
			if idx, ok := numMap[target]; ok && idx > j {
				arr := []int{nums[i], nums[j], target}
				res = append(res, arr)
			}
		}
	}
	return res
}
