package main

func countSubstrings(s string) int64 {
	res := int64(0)
	for i, v := range s {
		if v == '0' {
			continue
		}
		if v == '1' || v == '2' || v == '5' {
			res += int64(i + 1)
			continue
		}
		tmp :=
		for j := 0; j <= i; j++ {

		}
	}
}
