package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		arr  []int
		Want [][]int
	}

	testGroup := []testCase{
		{[]int{-1, 0, 1, 2, -1, -1, -4, 3}, [][]int{}},
	}

	for i, _ := range testGroup {
		result := combine(6, 3)
		// if !compareSlice(result, v.Want) {
		// 	t.Fatalf("%d, v %v, expect '%v' but '%v'", i, v, v.Want, result)
		// }
		fmt.Println(i, "result", result)
	}
}

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
