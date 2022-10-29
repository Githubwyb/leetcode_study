package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		op    []string
		value []int
		Want  []int
	}

	testGroup := []*testCase{
		{
			[]string{"ProductOfNumbers", "add", "getProduct", "getProduct", "add", "add", "getProduct", "add", "getProduct", "add", "getProduct", "add", "getProduct", "getProduct", "add", "getProduct"},
			[]int{0, 7, 1, 1, 4, 5, 3, 4, 4, 3, 4, 8, 1, 6, 2, 3},
			[]int{0, 0, 7, 7, 0, 0, 140, 0, 560, 0, 240, 0, 8, 13440, 0, 48},
		},
		{
			[]string{"ProductOfNumbers", "add", "getProduct", "getProduct", "getProduct", "add", "add", "add"},
			[]int{0, 1, 1, 1, 1, 7, 6, 7},
			[]int{0, 0, 1, 1, 1, 0, 0, 0},
		},
		{
			[]string{"ProductOfNumbers", "add", "add", "add", "add", "add", "getProduct", "getProduct", "getProduct", "add", "getProduct"},
			[]int{0, 3, 0, 2, 5, 4, 2, 3, 4, 8, 2},
			[]int{0, 0, 0, 0, 0, 0, 20, 40, 0, 0, 32},
		},
	}

	var product ProductOfNumbers
	for i, v := range testGroup {
		for j := range v.op {
			switch v.op[j] {
			case "ProductOfNumbers":
				product = Constructor()
				break

			case "add":
				product.Add(v.value[j])
				break

			case "getProduct":
				result := product.GetProduct(v.value[j])
				if result != v.Want[j] {
					t.Fatalf("%d, GetProduct %v failed, expect %v but %v", i, v.value[j], v.Want[j], result)
				}
				break
			}
		}
		fmt.Println(i, "input", v.op, v.value)
	}
	var product1 ProductOfNumbers1
	for i, v := range testGroup {
		for j := range v.op {
			switch v.op[j] {
			case "ProductOfNumbers":
				product1 = Constructor1()
				break

			case "add":
				product1.Add(v.value[j])
				break

			case "getProduct":
				result := product1.GetProduct(v.value[j])
				if result != v.Want[j] {
					t.Fatalf("%d, GetProduct %v failed, expect %v but %v", i, v.value[j], v.Want[j], result)
				}
				break
			}
		}
		fmt.Println(i, "input", v.op, v.value)
	}
}
