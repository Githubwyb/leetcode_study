package main

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	count := 0
	for _, v := range s {
		if v == '1' {
			count++
		}
	}
	if count == 0 {
		return s
	}
	res := make([]byte, n)
	count--
	for i := range s[:n-1] {
		if count > 0 {
			res[i] = '1'
			count--
			continue
		}
		res[i] = '0'
	}
	res[n-1] = '1'
	return string(res)
}