package main

func checkDistances(s string, distance []int) bool {
	// 使用数组作为map存储，0到128
	showMap := make([]int, 128)
	for i := range showMap {
		showMap[i] = -1
	}
	for i := range s {
		index := s[i]-'a'
		if showMap[index] == -1 {
			showMap[index] = i
			continue
		}

		d := i-showMap[index]-1
		if distance[index] != d {
			return false
		}
	}
	return true
}