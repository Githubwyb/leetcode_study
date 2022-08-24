package main

import (
	"fmt"
	_ "fmt"
	"strconv"
	"strings"
)

type transationInfo struct {
	time   int
	amount int
	city   string
	hasOut bool
}

func invalidTransactions(transactions []string) []string {
	// 交易转成对应需要的类型，使用map进行储存，按照名字进行划分
	// 同名交易，如果合法，替换前一个名字
	//          如果不合法，输出两个交易
	var result []string
	var recordMap map[string]transationInfo
	for _, v := range transactions {
		arr := strings.Split(v, ",")
		name := arr[0]
		var info transationInfo
		tmp, _ := strconv.ParseInt(arr[1], 10, 64)
		info.time = int(tmp)
		tmp, _ = strconv.ParseInt(arr[2], 10, 64)
		info.amount = int(tmp)
		info.city = arr[3]

		lastInfo, ok := recordMap[name]
		if !ok {
			recordMap[name] = info
			continue
		}

		// 存在判断是否合法
		if info.amount > 1000 {
			result = append(result, fmt.Sprintf("%s,%d,%d,%s"))
		} lastInfo.city != info.city && info.time - lastInfo.time < 60 {

		}
	}
}
