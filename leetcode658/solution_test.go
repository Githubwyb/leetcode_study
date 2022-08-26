package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		arr  []int
		k    int
		x    int
		Want []int
	}

	testGroup := []testCase{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 10, 10, 10}, 1, 9, []int{10}},
		{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
	}

	for i, v := range testGroup {
		tmpArr := make([]int, len(v.arr))
		copy(tmpArr, v.arr)
		result := findClosestElements(tmpArr, v.k, v.x)
		if len(result) != len(v.Want) {
			t.Fatalf("%d, v.arr %v k %d x %d, expect '%v' but '%v'", i, v.arr, v.k, v.x, v.Want, result)
		}
		for index, value := range result {
			if value != v.Want[index] {
				t.Fatalf("%d, v.arr %v k %d x %d, expect '%v' but '%v'", i, v.arr, v.k, v.x, v.Want, result)
			}
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		tmpArr := make([]int, len(v.arr))
		copy(tmpArr, v.arr)
		result := findClosestElements1(tmpArr, v.k, v.x)
		if len(result) != len(v.Want) {
			t.Fatalf("%d, v.arr %v k %d x %d, expect '%v' but '%v'", i, v.arr, v.k, v.x, v.Want, result)
		}
		for index, value := range result {
			if value != v.Want[index] {
				t.Fatalf("%d, v.arr %v k %d x %d, expect '%v' but '%v'", i, v.arr, v.k, v.x, v.Want, result)
			}
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		tmpArr := make([]int, len(v.arr))
		copy(tmpArr, v.arr)
		result := findClosestElements2(tmpArr, v.k, v.x)
		if len(result) != len(v.Want) {
			t.Fatalf("%d, v.arr %v k %d x %d, expect '%v' but '%v'", i, v.arr, v.k, v.x, v.Want, result)
		}
		for index, value := range result {
			if value != v.Want[index] {
				t.Fatalf("%d, v.arr %v k %d x %d, expect '%v' but '%v'", i, v.arr, v.k, v.x, v.Want, result)
			}
		}
		fmt.Println(i, "result", result)
	}
}
