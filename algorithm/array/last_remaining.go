package array

//https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
//约瑟夫环问题
func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

//timeout
func lastRemaining1(n int, m int) int {
	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i] = i
	}
	index := 0 //被删除元素的后一个元素的索引，每次删除第m个元素，从此元素开始计数
	for len(list) > 1 {
		//(index+m-1)是本次被删除元素的索引，删除后，被删除元素的后一个元素的索引变成了(index+m-1)
		index = (index + m - 1) % len(list)
		list = append(list[:index], list[index+1:]...)
	}
	return list[0]
}
