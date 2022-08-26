package main

import (
	_ "fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	compare := 2 * x
	for i, _ := range arr {
		if i+k == len(arr) {
			return arr[i:]
		}
		if arr[i+k]+arr[i] >= compare {
			return arr[i : i+k]
		}
	}
	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements1(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}

func findClosestElements2(arr []int, k int, x int) []int {
	right := sort.SearchInts(arr, x)
	left := right - 1
	if left < 0 {
		return arr[:k]
	}
	for right-left < k+1 {
		if left < 0 {
			right++
		} else if right >= len(arr) || arr[left]+arr[right] >= x+x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
