package main

func canJump(nums []int) bool {
	l, r := 0, 0
	for r < len(nums) {
		r = max(r, l+nums[l])
		l += 1
		if r < l {
			return len(nums)-1 <= r
		}
	}
	return true
}
