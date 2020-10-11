//https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/submissions/
//空间复杂度时间复杂度均为O(1)
package stack

type MinStack struct {
	arr []int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		arr: make([]int, 0),
		min: 1<<63 - 1,
	}
}

func (this *MinStack) Push(x int) {
	this.arr = append(this.arr, x-this.min)
	if x < this.min {
		this.min = x
	}
}

func (this *MinStack) Pop() {
	if l := len(this.arr); l > 0 {
		top := this.arr[l-1]
		if top < 0 {
			this.min -= top
		}
		this.arr = this.arr[:l-1]
	}
}

func (this *MinStack) Top() int {
	if l := len(this.arr); l > 0 {
		top := this.arr[l-1]
		if top >= 0 {
			return top + this.min
		}
		return this.min
	}
	return -1 << 63
}

func (this *MinStack) Min() int {
	return this.min
}
