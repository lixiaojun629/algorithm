package bit

import "testing"

func TestCountBits(t *testing.T) {
	bits := CountBits(100)
	t.Log(bits)
}
