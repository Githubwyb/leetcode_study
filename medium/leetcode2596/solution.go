package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func checkValidGrid(grid [][]int) bool {
	type pos struct {
		x, y int
	}
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid)
	posList := make([]pos, n*n)
	for i, v := range grid {
		for j, p := range v {
			posList[p] = pos{
				x: j,
				y: i,
			}
		}
	}
	for i := 1; i < len(posList); i++ {
		xd, yd := abs(posList[i].x-posList[i-1].x), abs(posList[i].y-posList[i-1].y)
		if xd+yd != 3 || abs(xd-yd) != 1 {
			return false
		}
	}
	return true
}
