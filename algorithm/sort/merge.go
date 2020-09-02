package sort

//Merge 归并排序
func Merge(a []int) {
	l := len(a)
	mergeSort(a, 0, l-1)
}

//自顶向下的归并排序
func mergeSort(a []int, l, h int) {
	if l >= h {
		return
	}
	m := l + (h-l)>>2
	mergeSort(a, l, m)
	mergeSort(a, m+1, h)
	merge(a, l, m, h)
}

//MergeBu 自底向上的归并排序, 非递归
func MergeBu(a []int) {
	for size := 1; size < len(a); size += size {
		for lo := 0; lo < len(a)-size; lo += size + size {
			merge(a, lo, lo+size-1, min(lo+size+size-1, len(a)-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(a []int, l, m, h int) {
	tmpArr := make([]int, h-l+1)
	tmpLeft := 0
	left := l
	right := m + 1
	for left <= m && right <= h {
		if a[left] <= a[right] {
			tmpArr[tmpLeft] = a[left]
			left++
		} else {
			tmpArr[tmpLeft] = a[right]
			right++
		}
		tmpLeft++
	}

	for left <= m {
		tmpArr[tmpLeft] = a[left]
		left++
		tmpLeft++
	}

	for right <= h {
		tmpArr[tmpLeft] = a[right]
		right++
		tmpLeft++
	}

	for i := 0; i < len(tmpArr); i++ {
		a[l+i] = tmpArr[i]
	}
}
