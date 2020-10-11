package array

//https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	l := l1 + l2
	if l&1 == 1 {
		mid := findKth(nums1, nums2, l/2+1)
		return float64(mid)
	}
	mid1 := findKth(nums1, nums2, l/2)
	mid2 := findKth(nums1, nums2, l/2+1)
	return float64(mid1+mid2) / 2.0
}

//从两个排序数组中找到第k小的数字
//nums1[k/2-1] 与 nums2[k/2-1]相比，每次淘汰较小的 k/2 个数字
func findKth(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	length1, length2 := len(nums1), len(nums2)
	for {
		if index1 == length1 {
			return nums2[index2+k-1]
		}
		if index2 == length2 {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		nindex1 := min(index1+k/2-1, length1-1)
		nindex2 := min(index2+k/2-1, length2-1)
		if nums1[nindex1] <= nums2[nindex2] {
			k -= (nindex1 - index1 + 1)
			index1 = nindex1 + 1

		} else if nums1[nindex1] > nums2[nindex2] {
			k -= (nindex2 - index2 + 1)
			index2 = nindex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//归并
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	l := l1 + l2
	nums := make([]int, l)
	i, j := 0, 0
	for k := 0; k < l; k++ {
		if i == l1 {
			nums[k] = nums2[j]
			j++
		} else if j == l2 {
			nums[k] = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}
	if l&1 == 1 {
		return float64(nums[l/2])
	}
	return float64(nums[l/2-1]+nums[l/2]) / 2.0
}
