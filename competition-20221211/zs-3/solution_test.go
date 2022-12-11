package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		op    []string
		param [][]int
		Want  []interface{}
	}

	testGroup := []testCase{
		{
			[]string{"Allocator", "allocate", "allocate", "allocate", "free", "allocate", "allocate", "allocate", "free", "allocate", "free"},
			[][]int{{10}, {1, 1}, {1, 2}, {1, 3}, {2}, {3, 4}, {1, 1}, {1, 1}, {1}, {10, 2}, {7}},
			[]interface{}{nil, 0, 1, 2, 1, 3, 1, 6, 3, -1, 0},
		},
	}

	for i, v := range testGroup {
		var al Allocator
		check := make([]interface{}, len(v.Want))
		for i := range v.op {
			switch v.op[i] {
			case "Allocator":
				al = Constructor(v.param[i][0])
				check[i] = nil
			case "allocate":
				result := al.Allocate(v.param[i][0], v.param[i][1])
				check[i] = result
				if result != v.Want[i] {
					t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, check)
				}
			case "free":
				result := al.Free(v.param[i][0])
				check[i] = result
				if result != v.Want[i] {
					t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, check)
				}
			}

		}
		fmt.Println(i, "result", check)
	}
}
