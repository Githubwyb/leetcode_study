package main

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	check := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0, 1)
	for _, v := range s {
		if t := check[byte(v)]; t > 0 {
			if len(stack) == 0 || t != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(v))
		}
	}
	return len(stack) == 0
}
