package main

func closestNodes(root *TreeNode, queries []int) [][]int {
	result := make([][]int, 0, len(queries))
	for _, v := range queries {
		min := -1
		max := -1
		point := root
		for point != nil {
			if point.Val == v {
				min, max = point.Val, point.Val
				break
			} else if point.Val < v {
				min = point.Val
				point = point.Right
			} else if point.Val > v {
				max = point.Val
				point = point.Left
			}
		}
		result = append(result, []int{min, max})
	}
	return result
}
