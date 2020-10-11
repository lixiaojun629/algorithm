package queue

//https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/
type CQueue struct {
	stack1 *Stack
	stack2 *Stack
}

func Constructor() CQueue {
	return CQueue{
		stack1: &Stack{list: make([]int, 0)},
		stack2: &Stack{list: make([]int, 0)},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Size() == 0 {
		for this.stack1.Size() > 0 {
			this.stack2.Push(this.stack1.Pop())
		}
	}

	if this.stack2.Size() > 0 {
		return this.stack2.Pop()
	}
	return -1
}

type Stack struct {
	list []int
}

func (s *Stack) Push(value int) {
	s.list = append(s.list, value)
}

func (s *Stack) Pop() int {
	if len(s.list) > 0 {
		val := s.list[len(s.list)-1]
		s.list = s.list[:(len(s.list) - 1)]
		return val
	}
	return -1
}

func (s *Stack) Size() int {
	return len(s.list)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
