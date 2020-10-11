package dp

//https://leetcode-cn.com/problems/maximum-product-subarray/
//nums[i]>0: dp[i][0] = max(dp[i-1][0]*nums[i],nums[i])
//           dp[i][1] = dp[i-1][1]*nums[i]
//nums[i]<0: dp[i][0] = dp[i-1][1]*nums[i]
//           dp[i][1] = min(dp[i-1][0]*nums[i],nums[i])

func MaxProduct(nums []int) int {
	//dp[i][0] 以i结尾的子数组的最大正乘积
	//dp[i][1] 以i结尾的子数组的最大负的乘积
	if len(nums) == 0 {
		return 0
	}

	var dp = make([][2]int, 2)
	var res int
	//正数组里面有负值，并不影响后面的正最大值的计算。
	dp[0][0], dp[0][1], res = nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		x, y := i%2, (i-1)%2
		if nums[i] > 0 {
			dp[x][0] = maxInt(nums[i]*dp[y][0], nums[i])
			dp[x][1] = nums[i] * dp[y][1]
		} else if nums[i] < 0 {
			dp[x][0] = nums[i] * dp[y][1]
			dp[x][1] = minInt(nums[i]*dp[y][0], nums[i])
		} else {
			dp[x][0] = 0
			dp[x][1] = 0
		}
		if dp[x][0] > res {
			res = dp[x][0]
		}
	}
	return res
}
