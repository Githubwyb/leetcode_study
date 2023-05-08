package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		root []interface{}
		Want []interface{}
	}

	testGroup := []testCase{
		{[]interface{}{7, 13, 11}, []interface{}{7, 11, 13}},
		{[]interface{}{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}, []interface{}{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1}},
	}

	for i, v := range testGroup {
		result := reverseOddLevels(MakeTreeNode(v.root))
		resultArr := Tree2Array(result)
		if !CompareSlice(resultArr, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, resultArr)
		}
		fmt.Println(i, "result", resultArr)
	}
	for i, v := range testGroup {
		result := reverseOddLevels1(MakeTreeNode(v.root))
		resultArr := Tree2Array(result)
		if !CompareSlice(resultArr, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, resultArr)
		}
		fmt.Println(i, "result", resultArr)
	}
	for i, v := range testGroup {
		result := reverseOddLevels2(MakeTreeNode(v.root))
		resultArr := Tree2Array(result)
		if !CompareSlice(resultArr, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, resultArr)
		}
		fmt.Println(i, "result", resultArr)
	}
}
