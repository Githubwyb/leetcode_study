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
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MDCXCV", 1695},
	}

	for i, v := range testGroup {
		result := romanToInt(v.s)
		if result != v.Want {
			t.Fatalf("%d, roman %v, expect %v but %v", i, v.s, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
	for i, v := range testGroup {
		result := romanToInt1(v.s)
		if result != v.Want {
			t.Fatalf("%d, roman %v, expect %v but %v", i, v.s, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
