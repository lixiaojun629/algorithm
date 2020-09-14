package recursive

//https://leetcode-cn.com/problems/powx-n/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	res := myPow(x, n>>1)
	res *= res
	if n&1 == 1 {
		res *= x
	}
	return res
}

func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return myPow(1/x, -n)
	}
	res := 1.0
	for n > 0 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
		n >>= 1
	}
	return res
}
