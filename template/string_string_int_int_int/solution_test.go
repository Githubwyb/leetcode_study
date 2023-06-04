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
		{"1", "12", 1, 8, 11},
	}

	for i, v := range testGroup {
		result := count(v.num1, v.num2, v.min_sum, v.max_sum)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
