package main

import (
	"sort"
)

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n := len(nums1)
	type pair struct{ x, y int }
	pairs := make([]pair, n)
	for i := range nums1 {
		pairs[i] = pair{
			x: nums1[i],
			y: nums2[i],
		}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].x > pairs[j].x
	})

	// 把索引放到第三位，这样动了queries不影响答案的结果
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] > queries[j][0]
	})

	var st []pair
	res := make([]int, len(queries))
	i := 0 // pairs的索引
	for _, v := range queries {
		x, y, index := v[0], v[1], v[2]
		for ; i < n && pairs[i].x >= x; i++ {
			p := pairs[i]
			if len(st) == 0 {
				st = []pair{p}
				continue
			}
			// 比最大的y大，才可以入栈
			if p.y <= st[len(st)-1].y {
				continue
			}
			// 从栈顶向栈底遍历，sum小于等于y的都不要了
			for i := len(st) - 1; i >= 0 && p.x+p.y >= st[i].x+st[i].y; i-- {
				st = st[:i]
			}
			st = append(st, p)
		}

		// 二分在栈内找第一个y大于等于query的即可
		i := sort.Search(len(st), func(i int) bool {
			return st[i].y >= y
		})
		if i == len(st) {
			res[index] = -1
		} else {
			res[index] = st[i].x + st[i].y
		}
	}
	return res
}

