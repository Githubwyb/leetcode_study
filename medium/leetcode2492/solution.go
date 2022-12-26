package main

import (
	"math"
)

type unionFind struct {
	parent []int
}

func initUnionFind(n int) unionFind {
	u := unionFind{}
	u.parent = make([]int, n)
	for i := range u.parent {
		u.parent[i] = i
	}
	return u
}

func (u unionFind) find(a int) int {
	ap := u.parent[a]
	// 找到最终节点
	for ap != u.parent[ap] {
		ap = u.parent[ap]
	}
	// 沿途都赋值最终节点
	for a != ap {
		u.parent[a], a = ap, u.parent[a]
	}
	return ap
}

func (u unionFind) merge(a, b int) {
	// b的父节点等于a的父节点，就是将两个点合并
	u.parent[u.find(b)] = u.find(a)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func minScore(n int, roads [][]int) int {
	union := initUnionFind(n + 1)
	for _, v := range roads {
		x, y := v[0], v[1]
		// 建立关系
		union.merge(x, y)
	}
	result := math.MaxInt
	for _, item := range roads {
		x, v := item[0], item[2]
		// 如果点和第一个点有关系代表能到达，那么就看路径是不是最小
		if union.find(x) == union.find(1) {
			result = min(result, v)
		}
	}
	return result
}
