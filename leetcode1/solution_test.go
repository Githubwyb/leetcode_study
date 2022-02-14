package leetcode1

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type testCase struct {
		InputArr *[]int
		Target   int
		Want     [2]int
	}

	testGroup := []*testCase{
		{&[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{&[]int{3, 2, 4}, 6, [2]int{1, 2}},
		{&[]int{3, 3}, 6, [2]int{0, 1}},
	}

	for i, v := range testGroup {
		result := TwoSum(*v.InputArr, v.Target)
		if result[0] != v.Want[0] || result[1] != v.Want[1] {
			t.Fatalf("%d input %v, target %d, expect %v but %v", i, v.InputArr, v.Target, v.Want, result)
		}
		fmt.Println(i, result)
	}
}
