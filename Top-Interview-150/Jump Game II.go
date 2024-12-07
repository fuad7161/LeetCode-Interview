package main

func jump(nums []int) int {
	dp := []int{}
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 100000)
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(nums) {
				dp[i+j] = min(dp[i]+1, dp[i+j])
			}
		}
	}
	return dp[len(nums)-1]
}
