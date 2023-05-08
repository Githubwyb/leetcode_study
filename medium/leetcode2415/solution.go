package main

import . "leetcode/common"

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

// bfs
func reverseOddLevels1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := []*TreeNode{root}
	for isOdd := 1; q[0].Left != nil; isOdd ^= 1 {
		// 将下一层放到一个新的数组中
		next := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			next = append(next, node.Left, node.Right)
		}
		q = next

		if isOdd == 0 {
			continue
		}
		// 如果是奇数层，将这一层反转值
		for i, n := 0, len(q); i < n/2; i++ {
			q[i].Val, q[n-1-i].Val = q[n-1-i].Val, q[i].Val
		}
	}
	return root
}

// dfs
func dfs(node1 *TreeNode, node2 *TreeNode, isOdd bool) {
	if node1 == nil {
		return
	}
	if isOdd {
		node1.Val, node2.Val = node2.Val, node1.Val
	}
	dfs(node1.Left, node2.Right, !isOdd)
	dfs(node1.Right, node2.Left, !isOdd)
}

func reverseOddLevels2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	dfs(root.Left, root.Right, true)
	return root
}
