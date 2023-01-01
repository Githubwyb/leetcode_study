package common

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTreeNode(input []interface{}) (root TreeNode) {
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
