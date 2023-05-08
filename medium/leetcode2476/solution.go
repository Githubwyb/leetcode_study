package main

import (
	. "leetcode/common"
	"sort"
)

var arr []int = make([]int, 1e5)

func closestNodes(root *TreeNode, queries []int) [][]int {
	arr = arr[:0]
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r.Left != nil {
			dfs(r.Left)
		}
		arr = append(arr, r.Val)
		if r.Right != nil {
			dfs(r.Right)
		}
	}
	dfs(root)

	n := len(arr)
	result := make([][]int, len(queries))
	for j, v := range queries {
		min, max := -1, -1
		i := sort.SearchInts(arr, v)
		if i == n {
			min = arr[i-1]
		} else if arr[i] == v {
			min, max = v, v
		} else if i == 0 {
			max = arr[i]
		} else {
			min, max = arr[i-1], arr[i]
		}
		result[j] = []int{min, max}
	}
	return result
}
