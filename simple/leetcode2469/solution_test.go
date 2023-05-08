package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		celsius float64
		Want    []float64
	}

	testGroup := []testCase{
		{36.50, []float64{309.65000, 97.70000}},
		{122.11, []float64{395.26000, 251.79800}},
	}

	for i, v := range testGroup {
		result := convertTemperature(v.celsius)
		if !common.CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
