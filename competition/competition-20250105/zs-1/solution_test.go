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
			nums: []int{1, 2, 1, 2, 1, 1, 1},
			Want: 5, // 根据所给示例，计算结果为5
		},
		{
			nums: []int{2, 3, 4, 5, 6},
			Want: 3, // 根据所给示例，计算结果为3
		},
		{
			nums: []int{1, 2, 3, 1, 4, 5, 1},
			Want: 5, // 类似上面所给案例的测试结果，可能是直接参考你提供的第3个案例结果，或者对一组有共同模式的数的优化策略得到的相同结果。
		},
	}

	for i, v := range testGroup {
		result := maxLength(DeepCopy(v.nums))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
