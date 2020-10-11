package dp

//https://leetcode-cn.com/problems/UlBDOe/
//f[i][j] 对[0...i]叶子进行调整，并且叶子i变成j状态需要的操作次数
//f[i][0] = f[i-1][0] + isYellow
//f[i][1] = min(f[i-1][0],f[i-1][1]) + isRed
//f[i][2] = min(f[i-1][1],f[i-1][2]) + isYellow
func minimumOperations(leaves string) int {
	n := len(leaves)
	//f[i][j] i表示第i片叶子，j表示状态 0:第一部分的红色，1:第二部分的黄色，2:第三部分的红色
	f := make([][3]int, n)
	f[0][0] = 0
	f[0][1] = n + 1
	f[0][2] = n + 1
	f[1][2] = n + 1
	if leaves[0] == 'y' {
		f[0][0] = 1
	}
	for i := 1; i < n; i++ {
		isYellow := 0
		if leaves[i] == 'y' {
			isYellow = 1
		}
		isRed := 0
		if leaves[i] == 'r' {
			isRed = 1
		}

		f[i][0] = f[i-1][0] + isYellow
		f[i][1] = min(f[i-1][1], f[i-1][0]) + isRed
		if i >= 2 {
			f[i][2] = min(f[i-1][2], f[i-1][1]) + isYellow
		}
	}
	return f[n-1][2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
