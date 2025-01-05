package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		grid [][]int
		k    int
		Want int64
	}

	testGroup := []testCase{
		{
			[][]int{
				{20, 27, 18},
				{37, 40, 19},
				{11, 14, 18},
				{8, 10, 9},
				{28, 32, 14},
				{1, 7, 5},
			},
			32,
			380, // 注意：将整数转为int64类型，以匹配结构体中的类型定义
		},
		{[][]int{{8, 10, 1}, {1, 3, 2}, {5, 6, 4}}, 4, 10},
		{
			[][]int{{23, 25, 3}, {34, 41, 4}, {5, 10, 16}},
			33,
			121,
		},
		{[][]int{{1, 10, 3}}, 2, 6},
	}

	for i, v := range testGroup {
		result := maximumCoins(DeepCopyIntss(v.grid), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
