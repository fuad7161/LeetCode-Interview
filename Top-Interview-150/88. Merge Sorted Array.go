/*
Explanation:
Appending: In Go, we use append() to add elements from nums2 into nums1. Since nums1 has a length of m, we use nums1[:m] to slice up to the m-th element and then append the nums2 array.
Sorting: The sort.Ints() function in Go is used to sort the integer slice nums1.
*/
package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println(n)
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
