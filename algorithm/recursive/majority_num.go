package recursive

//https://leetcode-cn.com/problems/majority-element/

func majorityElement(nums []int) int {
	return majorityElementCore(nums, 0, len(nums)-1)
}

func majorityElementCore(nums []int, low, high int) int {
	if low == high {
		return nums[low]
	}

	mid := low + (high-low)>>1
	left := majorityElementCore(nums, low, mid)
	right := majorityElementCore(nums, mid+1, high)
	if left == right {
		return left
	}
	leftCount := countNum(nums[low:mid+1], left)
	rightCount := countNum(nums[mid+1:high+1], right)
	if leftCount > rightCount {
		return left
	}
	return right
}

func countNum(nums []int, num int) int {
	count := 0
	for _, item := range nums {
		if item == num {
			count++
		}
	}
	return count
}

//摩尔投票法
func majorityElement1(nums []int) int {
	candidate := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
