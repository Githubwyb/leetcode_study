package main

func minimumPartition(s string, k int) int {
	t := 0
	result := 1
	for _, v := range s {
		d := int(byte(v) - '0')
		if d > k {
			return -1
		}
		tmp := t*10 + d
		if tmp > k {
			t = d
			result++
		} else {
			t = tmp
		}
	}
	return result
}
