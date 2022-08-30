package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid []string
		Want int
	}

	testGroup := []testCase{
		{[]string{"@.a.#", "###.#", "b.A.B"}, 8},
		{[]string{"@..aA", "..B#.", "....b"}, 6},
		{[]string{"@Aa"}, -1},
		{[]string{"@...a", ".###A", "b.BCc"}, 10},
	}

	for i, v := range testGroup {
		result := shortestPathAllKeys(v.grid)
		if result != v.Want {
			t.Fatalf("%d, v.grid %v expect '%v' but '%v'", i, v.grid, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := shortestPathAllKeys1(v.grid)
		if result != v.Want {
			t.Fatalf("%d, v.grid %v expect '%v' but '%v'", i, v.grid, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := shortestPathAllKeys2(v.grid)
		if result != v.Want {
			t.Fatalf("%d, v.grid %v expect '%v' but '%v'", i, v.grid, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
