package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		root []interface{}
		Want bool
	}

	testGroup := []testCase{
		{[]interface{}{1, 2, 2, nil, 3, nil, 3}, false},
		{[]interface{}{1, 2, 2, 3, 4, 4, 3}, true},
	}

	for i, v := range testGroup {
		result := isSymmetric(MakeTreeNode(v.root))
		if result != v.Want {
			t.Fatalf("%d, root %v, expect %v but %v", i, v.root, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}

	for i, v := range testGroup {
		result := isSymmetric1(MakeTreeNode(v.root))
		if result != v.Want {
			t.Fatalf("%d, root %v, expect %v but %v", i, v.root, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
