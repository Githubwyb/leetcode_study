package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int64
	}

	testGroup := []testCase{
		{[]int{7, 52, 2 ,4}, 596},
	}

	for i, v := range testGroup {
		result := findTheArrayConcVal(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
