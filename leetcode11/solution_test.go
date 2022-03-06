package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		height []int
		Want   int
	}

	testGroup := []*testCase{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for i, v := range testGroup {
		result := maxArea(v.height)
		if result != v.Want {
			t.Fatalf("%d, height %v, expect %v but %v", i, v.height, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
	for i, v := range testGroup {
		result := maxArea1(v.height)
		if result != v.Want {
			t.Fatalf("%d, height %v, expect %v but %v", i, v.height, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
