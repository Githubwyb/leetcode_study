package main

func specialPerm(nums []int) int {
	var mod int = 1e9 + 7
	// 先找出谁能和谁做邻居
	// key为数字值，v为index
	pMap := map[int][]int{}
	n := len(nums)
	for i, v := range nums {
		for j := i + 1; j < n; j++ {
			v1 := nums[j]
			if v%v1 == 0 || v1%v == 0 {
				pMap[v] = append(pMap[v], j)
				pMap[v1] = append(pMap[v1], i)
			}
		}
	}
	stMap := map[int]map[uint16]int{}
	for i, v := range nums {
		stMap[v] = map[uint16]int{
			1 << i: 1,
		}
	}
	for range nums[:n-1] {
		tmp := stMap
		stMap = make(map[int]map[uint16]int)
		// 遍历map，找state中没出现过的数字，看能否填到此位，能填就加上数量
		for k, v := range tmp {
			for _, v1 := range pMap[k] {
				for st, c := range v {
					if st&(1<<v1) > 0 {
						continue
					}
					st |= 1 << v1
					if stMap[nums[v1]] == nil {
						stMap[nums[v1]] = make(map[uint16]int)
					}
					stMap[nums[v1]][st] = (stMap[nums[v1]][st] + c) % mod
				}
			}
		}
	}
	res := 0
	for _, c := range stMap {
		for _, c1 := range c {
			res = (res + c1) % mod
		}
	}
	return res
}

func specialPerm1(nums []int) int {
	var mod int = 1e9 + 7
	n := len(nums)
	// 第一维是上一位是什么数字，第二维是状态（保存什么数字出现过的二进制），value为数量
	stMap := map[int]map[uint16]int{}
	for i, v := range nums {
		stMap[v] = map[uint16]int{
			1 << i: 1,
		}
	}
	for i := 1; i < n; i++ {
		tmp := stMap
		stMap = make(map[int]map[uint16]int)
		// 遍历map，找state中没出现过的数字，看能否填到此位，能填就加上数量
		for l, vMap := range tmp {
			// 遍历所有的状态，然后遍历所有的nums，没出现过且可以放就放进去
			for j, v := range nums {
				for st, c := range vMap {
					// 出现过或不能放到下一位就跳过
					if st&(1<<j) > 0 || (v%l != 0 && l%v != 0) {
						continue
					}
					if stMap[v] == nil {
						stMap[v] = make(map[uint16]int)
					}
					st |= (1 << j)
					stMap[v][st] = (stMap[v][st] + c) % mod
				}
			}
		}
	}
	res := 0
	for _, c := range stMap {
		for _, c1 := range c {
			res = (res + c1) % mod
		}
	}
	return res
}
