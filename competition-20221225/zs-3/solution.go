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
	acopy := a
	for a != u.parent[a] {
		a = u.parent[a]
	}
	for acopy != a {
		u.parent[acopy], acopy = a, u.parent[acopy]
	}
	return a
}

func (u unionFind) merge(a, b int) {
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
		union.merge(x, y)
	}
	result := math.MaxInt
	for _, item := range roads {
		x, v := item[0], item[2]
		if union.find(x) == union.find(1) {
			result = min(result, v)
		}
	}
	return result
}
