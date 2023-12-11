package main

func countTestedDevices(batteryPercentages []int) int {
	count := 0
	for _, v := range batteryPercentages {
		if v-count > 0 {
			count++
		}
	}
	return count
}
