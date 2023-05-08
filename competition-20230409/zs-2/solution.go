package main

func distance(nums []int) []int64 {
	res := make([]int64, len(nums))
	n := len(nums)
	// key均为值
	inMap := make(map[int]int)  // 正序的索引
	cntMap := make(map[int]int) // 正序的数量
	// 正序遍历
	for i, v := range nums {
		if j, ok := inMap[v]; ok {
			// 前面有
			res[i] = res[j] + int64(cntMap[v]*(i-j))
		}
		cntMap[v]++
		inMap[v] = i
	}
	inMap = make(map[int]int)     // 倒序的索引
	cntMap = make(map[int]int)    // 倒序的数量
	sumMap := make(map[int]int64) // 倒序的和
	for i := range nums {
		i = n - 1 - i
		v := nums[i]
		if j, ok := inMap[v]; ok {
			// 前面有
			sumMap[v] += int64(cntMap[v] * (j - i))
			res[i] += sumMap[v]
		}
		cntMap[v]++
		inMap[v] = i
	}
	return res
}

func distance1(nums []int) []int64 {
	grp := make(map[int][]int) // 数组存放索引
	for i, v := range nums {
		grp[v] = append(grp[v], i)
	}

	res := make([]int64, len(nums))
	// 遍历分组计算结果
	for _, v := range grp {
		n := len(v)
		if n == 1 {
			continue
		}
		s := make([]int, n) // 计算前缀和
		s[0] = v[0]
		for i := 1; i < n; i++ {
			s[i] = s[i-1] + v[i]
		}
		for i := range v {
			res[v[i]] = int64((i+1)*v[i] - s[i] + (s[n-1] - s[i]) - (n-1-i)*v[i])
		}
	}
	return res
}

func distance2(nums []int) []int64 {
	grp := make(map[int][]int) // 数组存放索引
	for i, v := range nums {
		grp[v] = append(grp[v], i)
	}

	res := make([]int64, len(nums))
	// 遍历分组计算结果
	for _, v := range grp {
		n := len(v)
		if n == 1 {
			continue
		}
		var s int64 = 0
		for _, v1 := range v[1:] {
			s += int64(v1 - v[0])
		}
		res[v[0]] = s
		for i := 1; i < n; i++ {
			res[v[i]] = res[v[i-1]] - int64((n-i-i)*(v[i]-v[i-1]))
		}
	}
	return res
}
