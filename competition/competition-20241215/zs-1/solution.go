package main

func buttonWithLongestTime(events [][]int) int {
	res, maxTime := events[0][0], events[0][1]
	for i, v := range events[1:] {
		k, diff := v[0], v[1]-events[i][1]
		if diff > maxTime || (diff == maxTime && res > k) {
			res, maxTime = k, diff
		}
	}
	return res
}
