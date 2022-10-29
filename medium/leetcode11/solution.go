package main

import (
	_ "fmt"
)

func maxArea(height []int) (result int) {
	left, right := 0, len(height)-1

	for right > left {
		minValue := height[left]
		if minValue > height[right] {
			minValue = height[right]
		}

		tmp := (right - left) * minValue
		if tmp > result {
			result = tmp
		}

		// less border move, move to a bigger border than minValue
		if height[left] < height[right] {
			for right > left && height[left] <= minValue {
				left++
			}
		} else {
			for right > left && height[right] <= minValue {
				right--
			}
		}
	}
	return
}

func maxArea1(height []int) (result int) {
	left, right := 0, len(height)-1

	for right > left {
		minValue := height[left]
		if minValue > height[right] {
			minValue = height[right]
		}
		tmp := (right - left) * minValue
		if result < tmp {
			result = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}
