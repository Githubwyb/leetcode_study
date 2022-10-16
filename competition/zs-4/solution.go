package main

func numberOfGoodPaths(vals []int, edges [][]int) int {
	// 从每个节点向外走，走过的不再走，看一下能走出几条路
	roads := make([][]bool, len(vals), len(vals))
	for i := range roads {
		roads[i] = make([]bool, len(vals), len(vals))
	}

	for _, v := range edges {
		x, y := v[0], v[1]
		roads[x][y] = true
		roads[y][x] = true
	}

	

}

func bfs()
