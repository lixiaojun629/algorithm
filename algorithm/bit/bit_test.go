package bit

import "testing"

func TestSingleNumber(t *testing.T) {
	singleNumbers([]int{1, 2, 5, 2})
}
func TestCountBits(t *testing.T) {
	bits := CountBits(100)
	t.Log(bits)
}
