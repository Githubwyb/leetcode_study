package main

import (
	"fmt"
	. "leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		i    string
		p1   [][]string
		r1   []float64
		p2   [][]string
		r2   []float64
		Want float64
	}

	testGroup := []testCase{
		{ // 第一个测试用例
			"EUR",
			[][]string{{"EUR", "USD"}, {"USD", "JPY"}},
			[]float64{2.0, 3.0},
			[][]string{{"JPY", "USD"}, {"USD", "CHF"}, {"CHF", "EUR"}},
			[]float64{4.0, 5.0, 6.0},
			720.0, // 计算结果应该为720.0（根据题目要求）
		},
		{ // 第二个测试用例
			"NGN",
			[][]string{{"NGN", "EUR"}},
			[]float64{9.0},
			[][]string{{"NGN", "EUR"}},
			[]float64{6.0},
			1.5, // 计算结果应该为1.5（根据题目要求）
		},
	}

	for i, v := range testGroup {
		result := maxAmount(v.i, DeepCopySlices(v.p1), DeepCopy(v.r1), DeepCopySlices(v.p2), DeepCopy(v.r2))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
