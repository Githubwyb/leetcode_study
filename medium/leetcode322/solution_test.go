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
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}

	for i, v := range testGroup {
		result := coinChange(v.coins, v.amount)
		if result != v.Want {
			t.Fatalf("%d, v.coins %v amount %d expect '%v' but '%v'", i, v.coins, v.amount, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := coinChange1(v.coins, v.amount)
		if result != v.Want {
			t.Fatalf("%d, v.coins %v amount %d expect '%v' but '%v'", i, v.coins, v.amount, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
