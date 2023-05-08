package leetcode

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		l1   []int
		l2   []int
		Want []int
	}

	testGroup := []*testCase{
		{[]int{0}, []int{0}, []int{0}},
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	for i, v := range testGroup {
		result := AddTwoNumbers(MakeList(v.l1), MakeList(v.l2))
		w := MakeList(v.Want)
		if !CompareList(result, w) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, result)
	}
}
