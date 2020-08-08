package sort

//Quick 快速排序
func Quick(a []int) {
	if len(a) <= 1 {
		return
	}
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, low, high int) {
	if low >= high {
		return
	}
	p := partition(a, low, high)
	quickSort(a, low, p-1)
	quickSort(a, p+1, high)
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	left, right := low, high-1
	for {
		for ; a[left] < pivot; left++ {
		}
		for ; a[right] > pivot; right-- {
		}
		if left < right {
			a[left], a[right] = a[right], a[left]
		} else {
			break
		}
	}
	a[left], a[high] = a[high], a[left]
	return left
}
