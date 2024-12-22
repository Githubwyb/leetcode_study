package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want int
	}

	testGroup := []testCase{
		{nums: []int{1, 2, 2, 3, 3, 4}, k: 2, Want: 6}, // 假设这里的6是函数计算后的结果
		{nums: []int{4, 4, 4, 4}, k: 1, Want: 3},       // 假设这里的3是函数计算后的结果
	}

	for i, v := range testGroup {
		result := maxDistinctElements(DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
