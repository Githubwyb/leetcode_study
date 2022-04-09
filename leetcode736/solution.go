package main

import (
	_ "fmt"
	"strconv"
)

// parseExpression, preprocess the expression
// convert '(let x 1 y 2 (let x 2 y (add 1 2) x))' to ['let' 'x' '1' 'y' '2' '(let x 2 y (add 1 2) x)']
func parseExpression(expression string) (result []string) {
	if expression[0] != '(' || expression[len(expression)-1] != ')' {
		panic("expression not valid")
	}

	for index := 1; index < len(expression)-1; index++ {
		// find full ()
		if expression[index] == '(' {
			stack := 1
			i := index + 1
			for ; i < len(expression); i++ {
				if expression[i] == '(' {
					stack++
					continue
				}

				if expression[i] == ')' {
					stack--
					if stack == 0 {
						break
					}
				}
			}

			result = append(result, expression[index:i+1])
			index = i + 1
			continue
		}

		// find next space
		i := index + 1
		for ; i < len(expression); i++ {
			if expression[i] == ' ' {
				break
			}
		}
		if i == len(expression) {
			result = append(result, expression[index:len(expression)-1])
			break
		}
		result = append(result, expression[index:i])
		index = i
	}
	return
}

// getValue, get value of input
// number => number
// var => paramMap[var]
// expression => parse(expression)
func getValue(input string, paramMap map[string]int) int {
	if input[0] == '(' {
		return parse(input, paramMap)
	}
	if v, ok := paramMap[input]; ok {
		return v
	}
	tmp, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(tmp)
}

// parse (let e1 v1 e2 v2 expr)
// will copy paramMap and make itself paramMap1
func parse(expression string, paramMap map[string]int) int {
	if expression[0] != '(' || expression[len(expression)-1] != ')' {
		panic("expression not valid")
	}
	paramMap1 := make(map[string]int)
	// copy paramMap
	for k, v := range paramMap {
		paramMap1[k] = v
	}

	// split expression to array
	arr := parseExpression(expression)
	if arr[0] == "add" || arr[0] == "mult" {
		if len(arr) != 3 {
			panic("add need three element, " + expression)
		}

		if arr[0] == "add" {
			return getValue(arr[1], paramMap1) + getValue(arr[2], paramMap1)
		}
		return getValue(arr[1], paramMap1) * getValue(arr[2], paramMap1)
	}

	if arr[0] == "let" {
		// make self scope
		for i := 2; i < len(arr)-1; i += 2 {
			paramMap1[arr[i-1]] = getValue(arr[i], paramMap1)
		}
		return getValue(arr[len(arr)-1], paramMap1)
	}
	panic("not support " + arr[0])
}

func evaluate(expression string) int {
	return parse(expression, nil)
}
