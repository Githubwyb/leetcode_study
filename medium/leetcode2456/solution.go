package main

type Author struct {
	id      string
	idPop   int // id对应的popular
	popular int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	allAuthor := make(map[string]Author)
	maxPopular := 0
	for i, name := range creators {
		id, view := ids[i], views[i]
		tmp := allAuthor[name]
		if tmp.id == "" {
			tmp = Author{
				id:      id,
				popular: view,
				idPop:   view,
			}
		} else {
			tmp.popular += view
			if view > tmp.idPop || (view == tmp.idPop && id < tmp.id) {
				tmp.id = id
				tmp.idPop = view
			}
		}
		if tmp.popular > maxPopular {
			maxPopular = tmp.popular
		}
		allAuthor[name] = tmp
	}

	result := make([][]string, 0)
	for k, v := range allAuthor {
		if v.popular == maxPopular {
			tmp := []string{k, v.id}
			result = append(result, tmp)
		}
	}

	return result
}
