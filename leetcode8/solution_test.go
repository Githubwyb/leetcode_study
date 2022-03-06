package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		Want int
	}

	testGroup := []*testCase{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"0032", 32},
		{"-2147483649", -2147483648},
		{"2147483648", 2147483647},
		{"+2147483648", 2147483647},
	}

	for i, v := range testGroup {
		result := myAtoi(v.s)
		if result != v.Want {
			t.Fatalf("%d, x %v expect '%v' but '%v'", i, v.s, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
