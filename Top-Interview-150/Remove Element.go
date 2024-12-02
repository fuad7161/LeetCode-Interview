/*
Explanation:
Two-Pointer Technique:

One pointer (index) tracks the position where the next non-val element should be placed.
The other pointer (_ in the for _, num loop) iterates through the array.
In-place Modification:

When an element not equal to val is found, it's assigned to nums[index], and index is incremented.
Efficiency:

The code operates in 𝑂(𝑛) time with O(1) additional space, as no auxiliary data structures are used.
*/
package main

func removeElement(nums []int, val int) int {
	index := 0
	for _, num := range nums {
		if num != val {
			nums[index] = num
			index++
		}
	}
	return index
}
