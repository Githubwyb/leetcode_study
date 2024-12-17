package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		Want int
	}

	testGroup := []testCase{
		{
			nums: []int{0, 0, 0, 3, 1, 3, 1, 3, 1, 1, 3, 1, 3, 3, 3, 0, 0, 0},
			Want: 26,
		},
		{
			nums: []int{0, 0, 2, 2, 0, 0, 0, 2, 2, 0, 0, 1},
			Want: 17,
		},
		{
			nums: []int{0, 0, 0, 2, 2, 0, 0, 0, 2, 2, 0, 0, 1},
			Want: 20,
		},
		{
			nums: []int{1, 2, 0, 1, 2, 1, 0, 0, 0, 2, 2, 0, 0, 0, 2, 2, 0, 0, 1}, // 输入值
			Want: 10,                                                             // 对应输出
		},
		{
			nums: []int{0, 0, 0, 0, 2, 2, 0, 1, 2, 1, 2}, // 输入值
			Want: 19,                                     // 对应输出
		},
		{
			nums: []int{3, 3, 3, 1, 3}, // 输入值
			Want: 3,                    // 对应输出
		},
		{
			nums: []int{1, 1, 0, 1, 3, 2, 2, 2}, // 输入值
			Want: 8,                             // 对应输出
		},
		{
			nums: []int{1, 1, 2, 1}, // 输入值
			Want: 2,                 // 对应输出
		},
		{
			nums: []int{1, 2, 3, 4}, // 输入值
			Want: 0,                 // 对应输出
		},
		{
			nums: []int{2, 3, 2, 2, 1}, // 输入值
			Want: 1,                    // 对应输出
		},
	}

	for i, v := range testGroup {
		result := beautifulSplits(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := beautifulSplits1(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
