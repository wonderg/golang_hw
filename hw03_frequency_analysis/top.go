package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	wordsStat := make(map[string]int)
	var words []string

	// Count words and fill the map and slice
	for _, word := range strings.Fields(input) {
		wordsStat[word]++
		if wordsStat[word] == 1 {
			words = append(words, word)
		}
	}

	// Sort slice based on counter in map
	sort.Slice(words, func(i, j int) bool {
		if wordsStat[words[i]] == wordsStat[words[j]] {
			return words[i] < words[j]
		}
		return wordsStat[words[i]] > wordsStat[words[j]]
	})

	top := 10
	if top > len(wordsStat) {
		top = len(wordsStat)
	}

	return words[:top]
}
