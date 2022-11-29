package main

import (
	"fmt"
	"testing"
)

func compareSlice(l, r interface{}) bool {
	switch l.(type) {
	case []int:
		a := l.([]int)
		b := r.([]int)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

	case []string:
		a := l.([]string)
		b := r.([]string)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

	case [][]string:
		a := l.([][]string)
		b := r.([][]string)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if !compareSlice(a[i], b[i]) {
				return false
			}
		}

	default:
		panic("not implement")
	}

	return true
}

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{[]int{1, 4, 4, 5, 5}, 4},
		{[]int{4, 4, 2, 4, 3}, 3},
		{[]int{1, 1, 1, 1, 1}, 0},
	}

	for i, v := range testGroup {
		result := unequalTriplets(v.nums)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
