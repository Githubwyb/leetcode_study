package main

func canMakePalindromeQueries(s string, queries [][]int) []bool {
	n := len(s)
	// 先明确s变成回文串需要改哪些位置，已经对应的不需要修改
	chArr := make([]int, 0, n)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			chArr = append(chArr, i)
		}
	}

	res := make([]bool, len(queries))
	for idx, q := range queries {
		if len(chArr) == 0 {
			res[idx] = true
			continue
		}
		a, b, c, d := q[0], q[1], q[2], q[3]
		lMap := map[byte]int{}
		rMap := map[byte]int{}
		can := true
		for _, i := range chArr {
			if !((i >= a && i <= b) || (i >= n-1-d && i <= n-1-c)) {
				// 两个都不在，明显不能处理
				can = false
				break
			}
			if i >= a && i <= b {
				lMap[s[i]]++
			}
			if i >= n-1-d && i <= n-1-c {
				rMap[s[n-1-i]]++
			}
		}
		if !can {
			continue
		}

		for _, i := range chArr {
			if i >= a && i <= b && i >= n-1-d && i <= n-1-c {
				// 重叠，先不处理
				continue
			}
			// 一半移动
			if i >= a && i <= b {
				ch := s[n-1-i]
				// 在左侧的map里，看右侧的字符在不在左侧map里
				if lMap[ch] == 0 {
					can = false
					break
				}
				if lMap[ch] == 1 {
					delete(lMap, ch)
				} else {
					lMap[ch]--
				}
				continue
			}
			ch := s[i]
			if rMap[ch] == 0 {
				can = false
				break
			}
			if rMap[ch] == 1 {
				delete(rMap, ch)
			} else {
				rMap[ch]--
			}
		}
		if !can || len(rMap) != len(lMap) {
			continue
		}
		// 重叠
		for k, v := range lMap {
			if v != rMap[k] {
				can = false
				break
			}
		}
		if !can {
			continue
		}
		res[idx] = true
	}

	return res
}

func canMakePalindromeQueries1(s string, queries [][]int) []bool {
	n := len(s)
	chArr := make([]int, 0, n)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			chArr = append(chArr, i)
		}
	}
	res := make([]bool, len(queries))
	for idx, q := range queries {
		if len(chArr) == 0 {
			res[idx] = true
			continue
		}
		a, b, c, d := q[0], q[1], q[2], q[3]
		lMap := map[byte]int{}
		rMap := map[byte]int{}
		for i := range s {
			if i >= a && i <= b {
				lMap[s[i]]++
			} else if i >= c && i <= d {
				rMap[s[i]]++
			}
		}
		can := true
		for i := 0; i < n/2; i++ {
			if !((i >= a && i <= b) || (i >= n-1-d && i <= n-1-c)) {
				// 非移动范围
				if s[i] != s[n-1-i] {
					can = false
					break
				}
				continue
			}
			if i >= a && i <= b && i >= n-1-d && i <= n-1-c {
				// 重叠，先不处理
				continue
			}
			// 一半移动
			if i >= a && i <= b {
				ch := s[n-1-i]
				// 在左侧的map里，看右侧的字符在不在左侧map里
				if lMap[ch] == 0 {
					can = false
					break
				}
				if lMap[ch] == 1 {
					delete(lMap, ch)
				} else {
					lMap[ch]--
				}
				continue
			}
			ch := s[i]
			if rMap[ch] == 0 {
				can = false
				break
			}
			if rMap[ch] == 1 {
				delete(rMap, ch)
			} else {
				rMap[ch]--
			}
		}
		if !can || len(rMap) != len(lMap) {
			continue
		}
		// 重叠
		for k, v := range lMap {
			if v != rMap[k] {
				can = false
				break
			}
		}
		if !can {
			continue
		}
		res[idx] = true
	}
	return res
}
