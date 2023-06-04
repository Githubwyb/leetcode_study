package main

func doesValidArrayExist(derived []int) bool {
	check := 0 // 第一个数默认为0
	for i := 0; i < len(derived); i++ {
		check ^= derived[i]
	}
	return check == 0
}
