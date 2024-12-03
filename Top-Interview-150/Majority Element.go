package main

func majorityElement(nums []int) int {
	mx, val := 0, 0
	for i := 0; i < len(nums); i++ {
		if mx == 0 {
			val = nums[i]
		}
		if val == nums[i] {
			mx += 1
		} else {
			mx -= 1
		}
	}
	return val
}

/*
solution is using map.
*/
func majorityElement1(nums []int) int {
	var vis = map[int]int{}
	mx := 0
	val := -1
	for i := 0; i < len(nums); i++ {
		vis[nums[i]] += 1
		if mx < vis[nums[i]] {
			mx = max(mx, vis[nums[i]])
			val = nums[i]
		}
	}
	return val
}
