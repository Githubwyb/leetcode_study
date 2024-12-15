package main

func findMaximumNumber(k int64, x int) int64 {
	sums := make([]int64, 1, 1) // 取每一个下标为1的总价值
	sumN := int64(0)
	for i := 1; ; i++ {
		var sum int64
		if i%x == 0 {
			// 这一位为1的价值是这一位为0的总价值加上这一位为1的总数量
			sum = sumN + 1<<(i-1)
		} else {
			sum = sumN << 1
		}
		if sum < k/2 {
			sums = append(sums, sum)
			sumN += sum
			continue
		}

	}
}
