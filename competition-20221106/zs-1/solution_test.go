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
		Want []int
	}

	testGroup := []testCase{
		{[]int{312, 312, 436, 892, 0, 0, 528, 0, 686, 516, 0, 0, 0, 0, 0, 445, 445, 445, 445, 445, 445, 984, 984, 984, 0, 0, 0, 0, 168, 0, 0, 647, 41, 203, 203, 241, 241, 0, 628, 628, 0, 875, 875, 0, 0, 0, 803, 803, 54, 54, 852, 0, 0, 0, 958, 195, 590, 300, 126, 0, 0, 523, 523}, []int{624, 436, 892, 528, 686, 516, 890, 890, 890, 1968, 984, 168, 647, 41, 406, 482, 1256, 1750, 1606, 108, 852, 958, 195, 590, 300, 126, 1046, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{[]int{1, 2, 2, 1, 1, 0}, []int{1, 4, 2, 0, 0, 0}},
		{[]int{0, 1}, []int{1, 0}},
	}

	for i, v := range testGroup {
		result := applyOperations(v.nums)
		if !compareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
