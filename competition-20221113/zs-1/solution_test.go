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
	case []float64:
		a := l.([]float64)
		b := r.([]float64)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] >= b[i] && (a[i]-b[i]) > 1e-5 {
				return false
			} else if a[i] < b[i] && (b[i]-a[i]) > 1e-5 {
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
		celsius float64
		Want    []float64
	}

	testGroup := []testCase{
		{36.50, []float64{309.65000, 97.70000}},
		{122.11, []float64{395.26000, 251.79800}},
	}

	for i, v := range testGroup {
		result := convertTemperature(v.celsius)
		if !compareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
