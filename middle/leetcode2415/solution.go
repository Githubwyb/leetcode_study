package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 转成广度遍历的数组
	nodeArr := []*TreeNode{root}
	for i := 0; i < len(nodeArr); i++ {
		tmp := nodeArr[i]
		if tmp.Left != nil {
			nodeArr = append(nodeArr, tmp.Left)
			tmp.Left = nil
		}
		if tmp.Right != nil {
			nodeArr = append(nodeArr, tmp.Right)
			tmp.Right = nil
		}
	}

	// 用数学计算遍历将对应的left和right全部替换完成
	layer := 0 // 层数
	layerItemSum := 1
	for index := 0; index < len(nodeArr); {
		nextBegin := index + layerItemSum
		if layer%2 == 1 {
			// 奇数层，倒序遍历添加孩子
			for i := layerItemSum - 1; i >= 0; i-- {
				if index+i >= len(nodeArr) {
					continue
				}
				// 下一层是偶数层，从头遍历
				leftIndex := nextBegin + (layerItemSum-1-i)*2
				rightIndex := leftIndex + 1
				if leftIndex < len(nodeArr) {
					nodeArr[index+i].Left = nodeArr[leftIndex]
				}
				if rightIndex < len(nodeArr) {
					nodeArr[index+i].Right = nodeArr[rightIndex]
				}

			}
		} else {
			// 偶数层，正序遍历添加孩子
			for i := 0; i < layerItemSum && index+i < len(nodeArr); i++ {
				// 下一层是奇数层，从尾部遍历
				leftIndex := nextBegin + layerItemSum*2 - 1 - i*2
				rightIndex := leftIndex - 1
				if leftIndex < len(nodeArr) {
					nodeArr[index+i].Left = nodeArr[leftIndex]
				}
				if rightIndex < len(nodeArr) {
					nodeArr[index+i].Right = nodeArr[rightIndex]
				}

			}
		}
		index += layerItemSum
		layer++
		layerItemSum *= 2
	}

	return root
}
