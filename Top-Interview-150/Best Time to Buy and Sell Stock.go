package main

func maxProfit(prices []int) int {
	ans, mn := 0, 1000000
	for i := 0; i < len(prices); i++ {
		mn = min(mn, prices[i])
		ans = max(ans, prices[i]-mn)
	}
	return ans
}
