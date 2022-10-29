package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		tokens []int
		power  int
		Want   int
	}

	testGroup := []testCase{
		{[]int{6, 0, 39, 52, 45, 49, 59, 68, 42, 37}, 99, 5},
		{[]int{}, 85, 0},
		{[]int{71, 55, 82}, 54, 0},
		{[]int{100}, 50, 0},
		{[]int{100, 200}, 150, 1},
		{[]int{100, 200, 300, 400}, 200, 2},
	}

	for i, v := range testGroup {
		result := bagOfTokensScore(v.tokens, v.power)
		if result != v.Want {
			t.Fatalf("%d, v.tokens %v, v.power %v, expect %v but %v", i, v.tokens, v.power, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := bagOfTokensScore1(v.tokens, v.power)
		if result != v.Want {
			t.Fatalf("%d, v.tokens %v, v.power %v, expect %v but %v", i, v.tokens, v.power, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
