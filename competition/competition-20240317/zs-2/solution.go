package main

func countSubstrings(s string, c byte) int64 {
	count := 0
	for _, v := range s {
		if byte(v) == c {
			count++
		}
	}

	return int64(count) * (1+int64(count)) / 2
}
