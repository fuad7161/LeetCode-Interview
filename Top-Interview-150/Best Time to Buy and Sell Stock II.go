package main

func maxProfit1(prices []int) int {
	mx, mn, ans := prices[0], prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			ans += mx - mn
			mx, mn = prices[i], prices[i]
		} else {
			mn = min(mn, prices[i])
			mx = max(mx, prices[i])
		}
	}
	return ans + mx - mn
}
