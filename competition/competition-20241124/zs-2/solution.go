package main

func isPossibleToRearrange(s string, t string, k int) bool {
	n := len(t)
	l := n / k
	tMap := map[string]int{}
	for i := 0; i < n; i += l {
		tMap[t[i:i+l]]++
	}

	for i := 0; i < n; i += l {
		if tMap[s[i:i+l]] == 0 {
			return false
		}
		tMap[s[i:i+l]]--
	}
	return true
}
