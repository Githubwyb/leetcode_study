package main

func minReverseOperations(n int, p int, banned []int, k int) []int {
	res := make([]int, n)
	for _, v := range banned {
		res[v] = -1
	}
	// bfs开始跳
	pq := make([]int, 0, n)
	tmp := make([]int, 0, n)
	pq = append(pq, p)
	for len(pq) > 0 {
		pq, tmp = tmp, pq
		pq = pq[:0]
		for _, v := range tmp {
			// 以此点为跳板开始跳
			// i为v在翻转数组的下标，翻转数组最左边为0，最右边为n-1
			i := k - n + v
			if i < 0 {
				i = 0
			}
			for ; i < k && i <= v; i++ {
				// i是v在k中的位置，k的0在v-i
				// 要翻转到的位置为0才能去
				t := k - 1 - i + v - i
				if t != v && res[t] == 0 {
					res[t] = res[v] + 1
					pq = append(pq, t)
				}
			}
		}
	}
	for i, v := range res {
		if v == 0 {
			res[i] = -1
		}
	}
	res[p] = 0
	return res
}
