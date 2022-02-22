package leetcode

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		Nums1 *[]int
		Nums2 *[]int
		Want  float64
	}

	testGroup := []*testCase{
		{&[]int{1, 3}, &[]int{2}, 2},
		{&[]int{1, 2}, &[]int{3, 4}, 2.5},
		{&[]int{100000}, &[]int{100001}, 100000.5},
		{&[]int{}, &[]int{2, 3}, 2.5},
		{&[]int{1}, &[]int{1}, 1},
		{&[]int{1}, &[]int{2, 3, 4}, 2.5},
		{&[]int{3}, &[]int{1, 2, 4}, 2.5},
		{&[]int{4}, &[]int{1, 2, 3}, 2.5},
		{&[]int{5, 6}, &[]int{1, 2, 3, 4}, 3.5},
	}

	for i, v := range testGroup {
		result := FindMedianSortedArrays(*v.Nums1, *v.Nums2)
		if result != v.Want {
			t.Fatalf("%d, nums1 %v, nums2 %v, expect %v but %v", i, v.Nums1, v.Nums2, v.Want, result)
		}
		fmt.Println(i, result)
	}

	for i, v := range testGroup {
		result := FindMedianSortedArrays1(*v.Nums1, *v.Nums2)		if result != v.Want {
			t.Fatalf("%d, nums1 %v, nums2 %v, expect %v but %v", i, v.Nums1, v.Nums2, v.Want, result)
		}
		fmt.Println(i, result)
	}

	for i, v := range testGroup {
		result := FindMedianSortedArrays2(*v.Nums1, *v.Nums2)
		if result != v.Want {
			t.Fatalf("%d, nums1 %v, nums2 %v, expect %v but %v", i, v.Nums1, v.Nums2, v.Want, result)
		}
		fmt.Println(i, result)
	}
}
