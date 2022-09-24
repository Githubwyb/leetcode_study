package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		number string
		Want   string
	}

	testGroup := []testCase{
		{"1-23-45 6", "123-456"},
		{"123 4-567", "123-45-67"},
		{"123 4-5678", "123-456-78"},
	}

	for i, v := range testGroup {
		result := reformatNumber(v.number)
		if result != v.Want {
			t.Fatalf("%d, v.number %v, expect '%v' but '%v'", i, v.number, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := reformatNumber1(v.number)
		if result != v.Want {
			t.Fatalf("%d, v.number %v, expect '%v' but '%v'", i, v.number, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
