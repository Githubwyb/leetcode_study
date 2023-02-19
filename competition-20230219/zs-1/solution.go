package main

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i, j := 0, 0
	n1, n2 := len(nums1), len(nums2)
	result := make([][]int, 0, n1+n2)
	for i != n1 && j != n2 {
		id1, v1, id2, v2 := nums1[i][0], nums1[i][1], nums2[j][0], nums2[j][1]
		if id1 == id2 {
			result = append(result, []int{id1, v1 + v2})
			i++
			j++
		} else if id1 < id2 {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	for ; i < n1; i++ {
		result = append(result, nums1[i])
		i++
	}
	for ; j < n2; j++ {
		result = append(result, nums2[j])
	}
	return result
}

type IntsSlice [][]int

func (x IntsSlice) Len() int { return len(x) }

// 为true，i向前；false，j向前。要满足相等时返回false
func (x IntsSlice) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x IntsSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func mergeArrays1(nums1 [][]int, nums2 [][]int) [][]int {
	intMap := make(map[int]int)
	for _, v := range nums1 {
		intMap[v[0]] += v[1]
	}
	for _, v := range nums2 {
		intMap[v[0]] += v[1]
	}
	result := make(IntsSlice, 0, len(intMap))
	for k, v := range intMap {
		result = append(result, []int{k, v})
	}
	sort.Sort(result)
	return result
}
