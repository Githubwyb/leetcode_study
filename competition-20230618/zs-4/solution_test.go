package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums     []int
		divisors []int
		Want     int
	}
	testGroup := []testCase{
		{[]int{937, 252, 716, 781, 319, 198, 273, 554, 140, 68, 694, 583, 1080, 16, 450, 229, 710, 1003, 1117, 1036, 398, 874, 289, 664, 600, 588, 372, 1066, 375, 532, 984, 328, 1067, 746}, []int{5, 3, 1, 3, 2, 1, 3, 3, 5, 3, 5, 5, 4, 1, 3, 1, 4, 4, 4, 1, 5, 1, 2, 3, 2, 3, 3, 4, 1, 3, 4, 1, 1, 5}, 1928},
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 2}, 3},
		{[]int{2, 3, 4, 2}, []int{1, 1, 1, 1}, 4},
	}

	for i, v := range testGroup {
		result := paintWalls(DeepCopy(v.nums), DeepCopy(v.divisors))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := paintWalls1(DeepCopy(v.nums), DeepCopy(v.divisors))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := paintWalls2(DeepCopy(v.nums), DeepCopy(v.divisors))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
