package main

func removeDuplicates1(nums []int) int {
	var vis = map[int]int{}
	j := 0
	for i := 0; i < len(nums); i++ {
		if vis[nums[i]] < 2 {
			nums[j] = nums[i]
			j += 1
		}
		vis[nums[i]] += 1
	}
	return j
}
