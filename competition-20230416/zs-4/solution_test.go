package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		n     int
		edges [][]int
		price []int
		trips [][]int
		Want  int
	}

	testGroup := []testCase{
		{2, [][]int{{0, 1}}, []int{2, 2}, [][]int{{0, 0}}, 1},
		{4, [][]int{{0, 1}, {1, 2}, {1, 3}}, []int{2, 2, 10, 6}, [][]int{{0, 3}, {2, 1}, {2, 3}}, 23},
		{40, [][]int{{0, 28}, {6, 29}, {7, 34}, {8, 5}, {5, 20}, {9, 12}, {12, 3}, {13, 11}, {14, 32}, {18, 3}, {3, 20}, {22, 15}, {15, 28}, {26, 25}, {25, 20}, {20, 17}, {27, 16}, {28, 2}, {31, 2}, {2, 21}, {21, 23}, {23, 4}, {4, 35}, {32, 19}, {33, 39}, {34, 10}, {10, 11}, {11, 16}, {16, 17}, {17, 1}, {1, 24}, {24, 30}, {30, 19}, {19, 39}, {35, 29}, {29, 38}, {36, 38}, {37, 39}, {38, 39}}, []int{4, 14, 4, 8, 26, 26, 12, 6, 10, 30, 30, 28, 2, 20, 8, 26, 10, 30, 18, 30, 18, 30, 16, 14, 18, 6, 20, 24, 20, 18, 8, 4, 12, 30, 12, 6, 30, 22, 28, 8}, [][]int{{10, 15}, {5, 21}, {16, 28}, {0, 31}, {13, 37}, {22, 27}, {13, 7}, {23, 10}, {7, 4}, {0, 11}, {35, 20}, {7, 12}, {16, 15}, {21, 6}, {7, 4}, {5, 25}, {10, 22}, {10, 1}, {20, 8}, {20, 23}, {38, 39}, {20, 2}}, 3041},
		{10, [][]int{{0, 4}, {2, 4}, {6, 3}, {8, 3}, {3, 4}, {4, 7}, {7, 5}, {5, 1}, {1, 9}}, []int{10, 6, 4, 8, 8, 6, 10, 8, 2, 2}, [][]int{{9, 6}, {5, 5}, {8, 4}, {7, 8}, {4, 5}, {0, 8}, {3, 4}, {6, 1}, {8, 0}, {4, 5}, {7, 5}, {9, 0}, {6, 3}, {0, 1}, {3, 9}, {6, 7}, {1, 5}, {0, 9}, {0, 4}, {2, 0}, {0, 3}, {1, 8}, {5, 3}, {6, 0}, {6, 4}, {9, 0}, {8, 7}, {5, 6}, {3, 6}, {8, 8}, {7, 8}, {6, 3}, {5, 7}, {5, 3}, {8, 7}, {7, 7}, {2, 5}, {4, 2}, {0, 8}, {3, 2}, {7, 2}, {1, 6}, {2, 7}, {1, 7}}, 826},
	}

	for i, v := range testGroup {
		result := minimumTotalPrice(v.n, common.DeepCopyIntss(v.edges), common.DeepCopy(v.price), common.DeepCopyIntss(v.trips))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := minimumTotalPrice1(v.n, common.DeepCopyIntss(v.edges), common.DeepCopy(v.price), common.DeepCopyIntss(v.trips))
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
