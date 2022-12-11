package main

import "fmt"

func sumOf3Equal0(a []int) [][]int {
	aMap := make(map[int]int) // key为元素，value为出现次数
	for i := range a {
		if _, ok := aMap[a[i]]; !ok {
			aMap[a[i]] = 0
		}
		aMap[a[i]]++
	}
	result := make([][]int, 0)
	for i := 0; i < len(a)-2; i++ {
		aMap[a[i]]--
		for j := i + 1; j < len(a)-1; j++ {
			aMap[a[j]]--
			check := a[i] + a[j]
			if v, ok := aMap[-check]; ok && v > 0 {
				for k := 0; k < v; k++ {
					result = append(result, []int{a[i], a[j], -check})
				}
			}

		}
		for j := i + 1; j < len(a)-1; j++ {
			aMap[a[j]]++
		}
	}
	return result
}

func combineI(i, j, k int, out []int, run func([]int)) {
	if k == 0 {
		run(out)
		return
	}
	for v := i; v <= j; v++ {
		out = append(out, v)
		combineI(v+1, j, k-1, out, run)
		out = out[:len(out)-1]
	}
}

func deepCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}

func getJ(i, j int) int {
	result := 1
	for v := i; v <= j; v++ {
		result *= v
	}
	return result
}

func combine(n, k int) [][]int {
	out := make([]int, 0, k)
	m := getJ(k+1, n) / getJ(1, n-k)
	fmt.Println(m)
	result := make([][]int, 0, m)
	combineI(1, n, k, out, func(outArr []int) {
		arr := make([]int, k)
		copy(arr, outArr)
		result = append(result, arr)
	})
	return result
}
