package helpers

import "strings"

func MinDistance(word1, word2 string) int {
	dp := make(map[[2]int]int)

	var getDist func(i, j int) int
	getDist = func(i, j int) int {
		if len(word1) == i && len(word2) == j {
			return 0
		}
		if len(word1) == i {
			return len(word2) - j
		}
		if len(word2) == j {
			return len(word1) - i
		}
		if val, ok := dp[[2]int{i, j}]; ok {
			return val
		}

		var result int
		if word1[i] == word2[j] || strings.EqualFold(string(word1[i]), string(word2[j])) {
			result = getDist(i+1, j+1)
		} else {
			result = 1 + min(getDist(i+1, j+1), min(getDist(i, j+1), getDist(i+1, j)))
		}

		dp[[2]int{i, j}] = result
		return result
	}

	return getDist(0, 0)
}
