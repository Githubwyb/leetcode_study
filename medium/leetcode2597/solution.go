package main

import "sort"

func beautifulSubsets(nums []int, k int) int {
	result := make(map[int]int)
	res := 0
	n := len(nums)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			res++
			return
		}
		// 不选
		dfs(i + 1)
		v := nums[i]
		if result[v-k] > 0 || result[v+k] > 0 {
			// 存在间隔为k的数
			return
		}
		// 可以选就选一下
		result[v]++
		dfs(i + 1)
		result[v]--
	}
	dfs(0)
	return res - 1 // 去除空集
}

func beautifulSubsets1(nums []int, k int) int {
	grp := make([]map[int]int, k)
	for _, v := range nums {
		i := v % k
		if grp[i] == nil {
			grp[i] = map[int]int{}
		}
		grp[i][v]++
	}

	res := 1
	type pair struct {
		x, cnt int
	}
	for _, v := range grp {
		if len(v) == 0 {
			continue
		}

		arr := make([]pair, 0, len(v))
		for k, val := range v {
			arr = append(arr, pair{x: k, cnt: val})
		}
		sort.Slice(arr, func(i, j int) bool { return arr[i].x < arr[j].x })

		ch, nch := 1<<arr[0].cnt-1, 1
		for i := 1; i < len(arr); i++ {
			ch1 := 1<<arr[i].cnt - 1
			if arr[i].x-k == arr[i-1].x {
				ch1 *= nch
			} else {
				ch1 *= (ch + nch)
			}
			nch = ch + nch
			ch = ch1
		}
		res *= (ch + nch)
	}
	return res - 1 // 去除空集
}
