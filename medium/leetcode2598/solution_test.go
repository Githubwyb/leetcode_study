package main

import (
	"fmt"
	"leetcode/common"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		nums []int
		k    int
		Want int
	}

	testGroup := []testCase{
		{[]int{-1000000000, -1320, -1319, -1318, -1317, -1316, -1315, -1314, -1313, -1312, -1311, -1310, -1309, -1308, -1307, -1306, -1305, -1304, -1303, -1302, -1301, -1300, -1299, -1298, -1297, -1296, -1295, -1294, -1293, -1292, -1291, -1290, -1289, -1288, -1287, -1286, -1285, -1284, -1283, -1282, -1281, -1280, -1279, -1278, -1277, -1276, -1275, -1274, -1273, -1272, -1271, -1270, -1269, -1268, -1267, -1266, -1265}, 132, 57},
		{[]int{3, 2, 3, 1, 0, 1, 4, 2, 3, 1, 4, 1, 3}, 5, 5},
		{[]int{1, -10, 7, 13, 6, 8}, 5, 4},
		{[]int{1, -10, 7, 13, 6, 8}, 7, 2},
	}

	for i, v := range testGroup {
		result := findSmallestInteger(common.DeepCopy(v.nums), v.k)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
