package common

import "math"

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	target := int(math.Sqrt(float64(num))) + 1
	for i := 2; i < target; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
