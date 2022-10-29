package main

import (
	_ "fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	result := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		index := 0
		for index+i < len(s) {
			// every line has index+i
			result = append(result, s[index+i])
			index += numRows + numRows - 2
			if i == 0 || i == numRows-1 {
				continue
			}

			// line except first line and last line, has index-i
			if index-i >= len(s) {
				break
			}
			result = append(result, s[index-i])
		}
	}
	return string(result)
}
