package sort

//https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(nums []int) int {
	tmp := make([]int, len(nums))
	return mergeSort1(nums, 0, len(nums)-1, tmp)
}

func mergeSort1(nums []int, low int, high int, tmp []int) int {
	if low >= high {
		return 0
	}
	mid := low + (high-low)>>1
	leftCount := mergeSort1(nums, low, mid, tmp)
	rightCount := mergeSort1(nums, mid+1, high, tmp)
	crossCount := mergeAndCount(nums, low, mid, high, tmp)
	return leftCount + rightCount + crossCount
}

func mergeAndCount(nums []int, low int, mid int, high int, tmp []int) int {
	i, j := low, mid+1
	count := 0
	for k := low; k <= high; k++ {
		if i == mid+1 {
			tmp[k] = nums[j]
			j++
		} else if j == high+1 {
			tmp[k] = nums[i]
			i++
		} else if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
			count += mid - i + 1
		}
	}
	for k := low; k <= high; k++ {
		nums[k] = tmp[k]
	}
	return count
}
