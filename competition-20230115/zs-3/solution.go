package main

func countGood(nums []int, k int) int64 {
	type arrInfo struct {
		s int
		e int
	}
	arrMap := make([]arrInfo, 0)
	// 查找到所有的对，按照尾部排序
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i] == nums[j] {
				arrMap = append(arrMap, arrInfo{
					s: j,
					e: i,
				})
			}
		}
	}

}
