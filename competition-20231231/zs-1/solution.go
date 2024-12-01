package main

func hasTrailingZeros(nums []int) bool {
	count := 0
	for _, v := range nums {
		if v&0x1 != 0 {
			continue
		}
		count++
		if count == 2 {
			return true
		}
	}
	return false
}
