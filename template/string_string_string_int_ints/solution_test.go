package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		s    string
		a    string
		b    string
		k    int
		Want []int
	}

	testGroup := []testCase{
		{"isawsquirrelnearmysquirrelhouseohmy", "my", "s", 15, []int{16, 33}},
	}

	for i, v := range testGroup {
		result := beautifulIndices(v.s, v.a, v.b, v.k)
		if !CompareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
