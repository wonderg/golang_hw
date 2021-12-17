package hw03frequencyanalysis

import (
	"log"
	"sort"
	"strings"
)

// WordCount holds word and count pair
type WordCount struct {
	word  string
	count int
}

func Top10(input string) []string {
	wordsStat := make(map[string]int)
	words := make([]string, 0)

	// Count words and fill the map and slice
	for _, word := range strings.Fields(input) {
		if _, ok := wordsStat[word]; ok {
			wordsStat[word]++
		} else {
			wordsStat[word] = 1
			words = append(words, word)
		}
	}

	log.Println("wordsStat:", wordsStat)
	log.Println("words non-sort: ", words)

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

	log.Println("words final: ", words[:top])
	return words[:top]
}
