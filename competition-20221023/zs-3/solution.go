package main

func minCost(nums []int, cost []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	minNum := nums[0]
	for _, v := range nums {
		if minNum > v {
			minNum = v
		}
	}

	var currentCost int64 = 0
	for i := range nums {
		currentCost += int64((nums[i] - minNum) * cost[i])
	}

	if currentCost == 0 {
		return currentCost
	}

	var nextCost int64
	for {
		nextCost = 0
		minNum++
		for i := range nums {
			if nums[i] >= minNum {
				nextCost -= int64(cost[i])
			} else {
				nextCost += int64(cost[i])
			}
		}
		if nextCost >= 0 {
			return currentCost
		}
		currentCost += nextCost
	}
}
