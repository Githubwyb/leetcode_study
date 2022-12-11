package main

import (
	"container/list"
	"fmt"
)

type test struct {
	head list.List
}

func main() {
	tmp := test{}
	tmp.head.PushFront(1)

	tmp2 := tmp
	tmp2.head.InsertAfter(2, tmp2.head.Front())

	for e := tmp2.head.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
