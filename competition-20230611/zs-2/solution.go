package main

func smallestString(s string) string {
	// 字典序最小，就是找第一个非a的到第一个a
	n := len(s)
	bs := []byte(s)
	for i, v := range bs {
		if v == 'a' {
			continue
		}
		for j := i; j < n && bs[j] != 'a'; j++ {
			bs[j]--
		}
		return string(bs)
	}
	// 全是a，把最后一个变成z
	return s[:n-1] + "z"
}
