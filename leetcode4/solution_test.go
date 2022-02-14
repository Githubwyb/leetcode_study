package leetcode2

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
		{&[]int{2, 3}, &[]int{2}, 2},
		{&[]int{1, 2}, &[]int{3, 4}, 2},
	}

	for i, v := range testGroup {
		result := FindMedianSortedArrays(*v.Nums1, *v.Nums2)
		if result != v.Want {
			t.Fatalf("%d, nums1 %v, nums2 %v, expect %v but %v", i, v.Nums1, v.Nums2, v.Want, result)
		}
		fmt.Println(i, result)
	}
}
