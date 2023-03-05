package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		root []interface{}
		k    int
		Want int64
	}

	testGroup := []testCase{
		{[]interface{}{5, 8, 9, 2, 1, 3, 7, 4, 6}, 2, 13},
		{[]interface{}{1, 2, nil, 3}, 1, 3},
	}

	for i, v := range testGroup {
		result := kthLargestLevelSum(MakeTreeNode(v.root), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}

func MakeTreeNode(input []interface{}) (root *TreeNode) {
	root = &TreeNode{}
	point := root
	pointChild := 0 // 0 self; 1 left; 2 right
	var treeList list.List
	for _, v := range input {
		if value, ok := v.(int); ok {
			switch pointChild {
			case 0:
				point.Val = value
				pointChild = 1
				break
			case 1:
				var tmp TreeNode
				tmp.Val = value
				point.Left = &tmp
				pointChild = 2
				treeList.PushBack(&tmp)
				break
			case 2:
				var tmp TreeNode
				tmp.Val = value
				point.Right = &tmp
				treeList.PushBack(&tmp)
				pointChild = 1
				point, _ = treeList.Front().Value.(*TreeNode)
				treeList.Remove(treeList.Front())
				break
			}
		} else {
			switch pointChild {
			case 0:
				pointChild = 1
				break
			case 1:
				point.Left = nil
				pointChild = 2
				treeList.PushBack(nil)
				break
			case 2:
				point.Right = nil
				treeList.PushBack(nil)
				pointChild = 1
				point = treeList.Front().Value.(*TreeNode)
				treeList.Remove(treeList.Front())
				break
			}
		}
	}
	return
}
