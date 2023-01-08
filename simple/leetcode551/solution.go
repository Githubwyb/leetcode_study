package main

func checkRecord(s string) bool {
	aCount := 0
	lCount := 0
	for i := range s {
		switch s[i] {
		case 'A':
			aCount++
			if aCount == 2 {
				return false
			}
			lCount = 0
		case 'L':
			lCount++
			if lCount == 3 {
				return false
			}
		default:
			lCount = 0
		}
	}
	return true
}
