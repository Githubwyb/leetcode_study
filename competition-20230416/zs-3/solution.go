package main

func addMinimum(word string) int {
	res := 0
	state := 0
	for i := 0; i < len(word); i++ {
		v := word[i]
		switch state {
		case 0:
			if v != 'a' {
				res++
				i--
			}
			state = 1
		case 1:
			if v != 'b' {
				res++
				i--
			}
			state = 2
		case 2:
			if v != 'c' {
				res++
				i--
			}
			state = 0
		}
	}
	if state != 0 {
		res += 3 - state
	}
	return res
}

func addMinimum1(word string) int {
	t := 1
	for i := 1; i < len(word); i++ {
		if word[i] <= word[i-1] {
			t++
		}
	}
	return 3*t - len(word)
}
