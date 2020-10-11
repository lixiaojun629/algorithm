package recursive

import "fmt"

func PrintNums(n int) {
	num := make([]uint8, n)
	print(num, 0, n)
}

func print(num []uint8, index int, max int) {
	if index == max {
		return
	}
	var i uint8
	for i = 0; i < 10; i++ {
		num[index] = i
		fmt.Println(num)
		print(num, index+1, max)
	}
}
