package main

import (
	_ "fmt"
)

var result []string

func getSubString(input string, output []byte, digitMap map[byte][]byte) {
	if len(input) == 0 {
		if len(output) != 0 {
			result = append(result, string(output))
		}
		return
	}
	numberCh := input[0]
	chList, ok := digitMap[numberCh]
	if !ok {
		getSubString(input[1:], output, digitMap)
		return
	}

	for _, ch := range chList {
		str := make([]byte, len(output), len(output)+1)
		copy(str, output)
		str = append(str, ch)
		getSubString(input[1:], str, digitMap)
	}
}

func getAllString(input string) []string {
	digitMap := make(map[byte][]byte)
	digitMap['2'] = []byte{'a', 'b', 'c'}
	digitMap['3'] = []byte{'d', 'e', 'f'}
	digitMap['4'] = []byte{'g', 'h', 'i'}
	digitMap['5'] = []byte{'j', 'k', 'l'}
	digitMap['6'] = []byte{'m', 'n', 'o'}
	digitMap['7'] = []byte{'p', 'q', 'r', 's'}
	digitMap['8'] = []byte{'t', 'u', 'v'}
	digitMap['9'] = []byte{'w', 'x', 'y', 'z'}
	result = make([]string, 0)

	output := []byte{}
	getSubString(input, output, digitMap)
	return result
}
