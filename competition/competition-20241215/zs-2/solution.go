package main

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
	// 先做两个汇率map
	getRateMap := func(ik string, pairs [][]string, rates []float64) map[string]float64 {
		type pair struct {
			to   string
			rate float64
		}
		res := map[string]float64{}
		m := make(map[string][]pair)
		for i, v := range pairs {
			a, b := v[0], v[1]
			m[a] = append(m[a], pair{b, rates[i]})
			m[b] = append(m[b], pair{a, 1 / rates[i]})
		}
		var dfs func(k string)
		dfs = func(k string) {
			r := res[k]
			for _, v := range m[k] {
				if _, ok := res[v.to]; ok {
					continue
				}
				res[v.to] = v.rate * r
				dfs(v.to)
			}
		}

		for _, v := range m[ik] {
			res[v.to] = v.rate
			dfs(v.to)
		}
		return res
	}
	rates1Map := getRateMap(initialCurrency, pairs1, rates1)
	rates2Map := getRateMap(initialCurrency, pairs2, rates2)
	maxRate := 1.0
	for k, v1 := range rates1Map {
		if v2, ok := rates2Map[k]; ok {
			if v1/v2 > maxRate {
				maxRate = v1 / v2
			}
		}
	}
	return maxRate
}
