package main

func removeDuplicates(nums []int) int {
	var vis = map[int]bool{}
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if vis[nums[i]] == false {
			nums[cnt] = nums[i]
			cnt += 1
		}
		vis[nums[i]] = true
	}
	return cnt

}
