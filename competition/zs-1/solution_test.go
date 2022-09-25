package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		names   []string
		heights []int
		Want    []string
	}

	testGroup := []testCase{
		{[]string{"Mary", "John", "Emma"}, []int{180, 165, 170}, []string{"Mary", "Emma", "John"}},
		{[]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}, []string{"Bob", "Alice", "Bob"}},
	}

	for i, v := range testGroup {
		result := sortPeople(v.names, v.heights)
		if len(result) != len(v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		for i := range result {
			if result[i] != v.Want[i] {
				t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
			}
		}
		fmt.Println(i, "result", result)
	}
}
