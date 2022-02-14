package leetcode2

import (
	_ "fmt"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	var l1, r1, l2, r2 int
	l1 = 0
	l2 = 0
	r1 = len(nums1) - 1
	r2 = len(nums2) - 1
	t1, t2 := (len(nums1)+len(nums2))/2, (len(nums1)+len(nums2)+1)/2
	for {
		m1 := (l1 + r1) / 2
		m2 := (l2 + r2) / 2
		if nums1[m1] > nums2[m2] {
			
		} else {

		}
	}
	return
}
