package main

import (
	_ "fmt"
	"strconv"
	"strings"
)

type transactionInfo struct {
	name      string
	time      int64
	amount    int64
	city      string
	hasOut    bool
	originStr string
}

func parseTransaction(str string) (info transactionInfo) {
	arr := strings.Split(str, ",")
	info.name = arr[0]
	info.time, _ = strconv.ParseInt(arr[1], 10, 64)
	info.amount, _ = strconv.ParseInt(arr[2], 10, 64)
	info.city = arr[3]
	info.hasOut = false
	info.originStr = str
	return
}

func appendResult(info *transactionInfo, result *[]string) {
	if info.hasOut {
		return
	}
	*result = append(*result, info.originStr)
	info.hasOut = true
}

func invalidTransactions(transactions []string) []string {
	recordMap := make(map[string][]transactionInfo)

	var result []string
	for _, v := range transactions {
		info := parseTransaction(v)

		// 先判断是否超出1000
		if info.amount > 1000 {
			appendResult(&info, &result)
		}

		// 判断是否之前有记录，不存在直接继续
		lastInfos, ok := recordMap[info.name]
		if !ok {
			recordMap[info.name] = []transactionInfo{info}
			continue
		}

		// 判断是否是小于60并不在同一个城市
		for i, lastInfo := range lastInfos {
			if lastInfo.city != info.city && (info.time-lastInfo.time <= 60 && info.time-lastInfo.time >= -60) {
				appendResult(&info, &result)
				appendResult(&lastInfos[i], &result)
			}
		}
		lastInfos = append(lastInfos, info)
		recordMap[info.name] = lastInfos
	}

	return result
}

func invalidTransactions1(transactions []string) []string {
	var result []string

	// 遍历当前交易，和后面冲突就记录，跳出继续
	for i, v := range transactions {
		info := parseTransaction(v)

		if info.amount > 1000 {
			result = append(result, info.originStr)
			continue
		}

		for j, v1 := range transactions {
			if j == i {
				continue
			}
			info1 := parseTransaction(v1)
			if info.name == info1.name && info1.city != info.city && (info.time-info1.time <= 60 && info.time-info1.time >= -60) {
				result = append(result, info.originStr)
				break
			}

		}
	}
	return result
}
