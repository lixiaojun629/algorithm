package sort


//Selection 选择排序，选择最小元素，和第一个元素交换位置，在剩下元素中选择最小元素，和第二个元素交换位置，依次到结束。
func Selection(arr []int){
	for i := 0 ; i < len(arr); i++ {
		min := i
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min],arr[i]
	}
}