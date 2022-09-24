package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		op     string
		values []int
		Want   int
	}

	testGroup := []*testCase{
		{"MyHashMap", []int{}, 0},
		{"put", []int{1, 1}, 0},
		{"put", []int{2, 2}, 0},
		{"get", []int{1}, 1},
		{"get", []int{3}, -1},
		{"put", []int{2, 1}, 0},
		{"get", []int{2}, 1},
		{"remove", []int{2}, 0},
		{"get", []int{2}, -1},
	}

	var hashMap MyHashMap
	for i, v := range testGroup {
		switch v.op {
		case "MyHashMap":
			hashMap = Constructor()
			break

		case "put":
			hashMap.Put(v.values[0], v.values[1])
			break

		case "get":
			result := hashMap.Get(v.values[0])
			if result != v.Want {
				t.Fatalf("%d, get %v failed, expect %v but %v", i, v.values[0], v.Want, result)
			}
			break

		case "remove":
			hashMap.Remove(v.values[0])
			break
		}
		fmt.Println(i, "input", v.op, v.values)
	}


	var hashMap1 MyHashMap1
	for i, v := range testGroup {
		switch v.op {
		case "MyHashMap":
			hashMap1 = Constructor1()
			break

		case "put":
			hashMap1.Put(v.values[0], v.values[1])
			break

		case "get":
			result := hashMap1.Get(v.values[0])
			if result != v.Want {
				t.Fatalf("%d, get %v failed, expect %v but %v", i, v.values[0], v.Want, result)
			}
			break

		case "remove":
			hashMap1.Remove(v.values[0])
			break
		}
		fmt.Println(i, "input", v.op, v.values)
	}
}
