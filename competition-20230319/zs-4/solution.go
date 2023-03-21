package main

import "math"

func findSmallestInteger(nums []int, value int) int {
	chooseMap := make([]int, value)
	for _, v := range nums {
		i := 0
		if v >= 0 {
			i = v % value
		} else {
			i = (value - (-v)%value) % value
		}
		chooseMap[i]++
	}
	min, index := math.MaxInt, -1
	for i, v := range chooseMap {
		if v == 0 {
			return i
		}
		if v < min {
			min = v
			index = i
		}
	}
	return index + min*value
}

func findSmallestInteger1(nums []int, value int) int {
	unChooseMap := make(map[int]int)
	max, min := math.MinInt, math.MaxInt
	for _, v := range nums {
		unChooseMap[v]++
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	for i := 0; i < len(nums); i++ {
		if unChooseMap[i] > 0 {
			unChooseMap[i]--
			continue
		}
		big, lit := i, i
		for {
			if big >= max && lit <= min {
				return i
			}
			if big < max {
				big += value
				if unChooseMap[big] > 0 {
					unChooseMap[big]--
					break
				}
			}
			if lit > min {
				lit -= value
				if unChooseMap[lit] > 0 {
					unChooseMap[lit]--
					break
				}
			}
		}
	}
	return len(nums)
}
