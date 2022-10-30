package main

func getAdd(n int64) ([]int, int) {
	result := 0
	resultArr := []int{}
	for n != 0 {
		add := int(n % 10)
		result += add
		resultArr = append(resultArr, add)
		n = n / 10
	}
	return resultArr, result
}

func makeIntegerBeautiful(n int64, target int) int64 {
	numArr, addResult := getAdd(n)
	if addResult <= target {
		return 0
	}

	var result int64 = 0
	diff := addResult - target
	aa := 1
	for _, v := range numArr {
		// 33 => 40 少了2 如果target是4或5都可以，diff为2或1
		if diff < v {
			result += int64(9-v) * int64(aa)
			return result+1
		}

		result += int64(9-v) * int64(aa)
		aa *= 10
		diff -= v
	}
	return result+1
}
