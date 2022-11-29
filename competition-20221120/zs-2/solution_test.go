package main

import (
	"container/list"
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		tree    []interface{}
		queries []int
		Want    [][]int
	}

	testGroup := []testCase{
		{[]interface{}{6, 2, 13, 1, 4, 9, 15, nil, nil, nil, nil, nil, nil, 14}, []int{2, 5, 16}, [][]int{{2, 2}, {4, 6}, {15, -1}}},
		{[]interface{}{4, nil, 9}, []int{3}, [][]int{{-1, 4}}},
	}

	for i, v := range testGroup {
		result := closestNodes(makeTreeNode(v.tree), v.queries)
		if !compareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}

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

	case [][]int:
		a := l.([][]int)
		b := r.([][]int)
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTreeNode(input []interface{}) (root *TreeNode) {
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
