package recursive

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func TestPrintNums(t *testing.T) {
	PrintNums(2)
}

func TestSlice(t *testing.T) {
	arr := []int{1}
	add(arr)
	fmt.Println(arr)

}
func add(arr []int) {
	arr = append(arr, 0)
	fmt.Println("add arr", arr)
}
