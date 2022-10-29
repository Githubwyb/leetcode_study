package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		num  []int
		Want int
	}

	testGroup := []testCase{
		{[]int{1, 3, 8, 48, 10}, 3},
		{[]int{3, 1, 5, 11, 13}, 1},
		{[]int{744437702, 379056602, 145555074, 392756761, 560864007, 934981918, 113312475, 1090, 16384, 33, 217313281, 117883195, 978927664}, 3},
	}

	for i, v := range testGroup {
		result := longestNiceSubarray(v.num)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}

	for i, v := range testGroup {
		result := longestNiceSubarray1(v.num)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
