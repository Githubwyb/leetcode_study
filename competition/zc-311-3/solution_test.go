package main

import (
	"container/list"
	"fmt"
	"testing"
)

func getTreeByArray(root []int) *TreeNode {
	if len(root) == 0 {
		return nil
	}

	result := &TreeNode{}
	var queue list.List
	var tmp *TreeNode
	for i := range root {
		if i == 0 {
			result.Val = root[0]
			queue.PushBack(result)
			continue
		}
		if tmp == nil {
			tmp = queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())
			tmp.Left = &TreeNode{}
			tmp.Left.Val = root[i]
			queue.PushBack(tmp.Left)
		} else {
			tmp.Right = &TreeNode{}
			tmp.Right.Val = root[i]
			queue.PushBack(tmp.Right)
			tmp = nil
		}
	}
	return result
}

func getArrayByTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	var queue list.List
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmp := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		result = append(result, tmp.Val)
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)

		}
		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
	return result
}

func TestSolution(t *testing.T) {
	type testCase struct {
		root []int
		Want []int
	}

	testGroup := []testCase{
		{[]int{7, 13, 11}, []int{7, 11, 13}},
		{[]int{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2}, []int{0, 2, 1, 0, 0, 0, 0, 2, 2, 2, 2, 1, 1, 1, 1}},
	}

	for i, v := range testGroup {
		result := reverseOddLevels(getTreeByArray(v.root))
		resultArr := getArrayByTree(result)
		if len(resultArr) != len(v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, resultArr)
		}
		for j := range resultArr {
			if resultArr[j] != v.Want[j] {
				t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, resultArr)
			}
		}
		fmt.Println(i, "result", resultArr)
	}
}
