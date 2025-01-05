package main

func calculateScore(s string) int64 {
	sc := int64(0)
	m := make([][]int64, 26)
	for i, v := range s {
		k := v - 'a'
		kj := 25 - k
		n := len(m[kj])
		if n == 0 {
			m[k] = append(m[k], int64(i))
			continue
		}
		sc += int64(i) - m[kj][n-1]
		m[kj] = m[kj][:n-1]
	}
	return sc
}
