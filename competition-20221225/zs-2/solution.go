package main

import "container/list"

type state struct {
	l     int
	r     int
	arr   []int
	count int
}

func check(arr []int, k int) bool {
	for _, v := range arr {
		if v < k {
			return false
		}
	}
	return true
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	seen := make([][]bool, len(s)+2)
	for i := range seen {
		seen[i] = make([]bool, len(s)+2)
	}
	queue := list.New()
	queue.PushBack(state{
		l:     0,
		r:     len(s) + 1,
		arr:   make([]int, 3),
		count: 0,
	})
	seen[0][len(s)+1] = true
	for queue.Len() > 0 {
		tmp := queue.Front().Value.(state)
		queue.Remove(queue.Front())
		result := tmp.count + 1
		// 左侧加
		for tmp.l < len(s) {
			l, r := tmp.l+1, tmp.r
			if l == r || seen[l][r] {
				break
			}
			seen[l][r] = true
			arr := make([]int, len(tmp.arr))
			copy(arr, tmp.arr)
			arr[int(s[l-1]-'a')]++
			if check(arr, k) {
				return result
			}
			queue.PushBack(state{
				l:     l,
				r:     r,
				arr:   arr,
				count: result,
			})
			break
		}

		// 右侧减
		for tmp.r > 1 {
			l, r := tmp.l, tmp.r-1
			if l == r || seen[l][r] {
				break
			}
			seen[l][r] = true
			arr := make([]int, len(tmp.arr))
			copy(arr, tmp.arr)
			arr[int(s[r-1]-'a')]++
			if check(arr, k) {
				return result
			}
			queue.PushBack(state{
				l:     l,
				r:     r,
				arr:   arr,
				count: result,
			})
			break
		}
	}
	return -1
}
