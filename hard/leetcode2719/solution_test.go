package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		num1    string
		num2    string
		min_sum int
		max_sum int
		Want    int
	}
	testGroup := []testCase{
		{"1", "12", 5, 8, 4},
		{"3103", "9832", 6, 7, 45},
		{"8990", "9927", 1, 3, 0},
		{"4179205230", "7748704426", 8, 46, 883045899},
		{"1", "12", 1, 8, 11},
		{"1", "5", 1, 5, 5},
	}

	for i, v := range testGroup {
		result := count(v.num1, v.num2, v.min_sum, v.max_sum)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := count1(v.num1, v.num2, v.min_sum, v.max_sum)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := count2(v.num1, v.num2, v.min_sum, v.max_sum)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
