package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		coins  []int
		amount int
		Want   int
	}

	testGroup := []testCase{
		{[]int{1, 2, 3, 4, 5}, 30, 4},
	}

	for i, v := range testGroup {
		result := getMinCount(v.coins, v.amount)
		if result != v.Want {
			t.Fatalf("%d, v.coins %v amount %d expect '%v' but '%v'", i, v.coins, v.amount, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
