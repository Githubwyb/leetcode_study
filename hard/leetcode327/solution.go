package main

type segTree []struct {
	l   int // 左端点
	r   int // 右端点
	val int // 值
}

func (t segTree) build() {

}

func countRangeSum(nums []int, lower int, upper int) int {
	// 计算前缀和，区间和就是两个端点相减，问题转化为区间值的个数
	n := len(nums)
	preSum := make([]int64, n)
	for i, v := range nums {
		if i == 0 {
			preSum[i] = int64(v)
			continue
		}
		preSum[i] = preSum[i-1] + int64(v)
	}

	return 0
}
