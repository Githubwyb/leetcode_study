package main

import (
	"container/list"
	"fmt"
	"testing"
)

func makeTreeNode(input []interface{}) (root TreeNode) {
	point := &root
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

func getTreeArray(root TreeNode) (input []interface{}) {
	var treeList list.List
	treeList.PushBack(&root)
	for e := treeList.Front(); e != nil; e = e.Next() {
		input = append(input, e.Value.(*TreeNode).Val)
		if e.Value.(*TreeNode).Left != nil {
			treeList.PushBack(e.Value.(*TreeNode).Left)
		}
		if e.Value.(*TreeNode).Right != nil {
			treeList.PushBack(e.Value.(*TreeNode).Right)
		}
	}
	return
}

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
		tmp := makeTreeNode(v.root)
		inputArr := getTreeArray(tmp)
		result := isSymmetric(&tmp)
		if result != v.Want {
			t.Fatalf("%d, root %v, expect %v but %v", i, v.root, v.Want, result)
		}
		fmt.Println(i, "inputArr", inputArr)
	}

	for i, v := range testGroup {
		tmp := makeTreeNode(v.root)
		inputArr := getTreeArray(tmp)
		result := isSymmetric1(&tmp)
		if result != v.Want {
			t.Fatalf("%d, root %v, expect %v but %v", i, v.root, v.Want, result)
		}
		fmt.Println(i, "inputArr", inputArr)
	}
}
