package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want bool
	}

	testGroup := []testCase{
		{[]int{4, 9, 6, 10}, true},
	}

	for i, v := range testGroup {
		result := primeSubOperation(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
