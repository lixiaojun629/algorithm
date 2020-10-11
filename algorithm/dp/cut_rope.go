package dp

func cuttingRope(n int) int {
	products := make([]int, n+1)
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	products[1] = 1
	products[2] = 2
	products[3] = 3
	if n >= 4 {
		for i := 4; i <= n; i++ {
			for j := 1; j <= i/2; j++ {
				products[i] = max(products[i], products[j]*products[i-j])
			}
		}
	}
	return products[n]
}
