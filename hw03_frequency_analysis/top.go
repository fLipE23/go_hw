package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Top10(input string) []string {
	regex := regexp.MustCompile(`[^A-Za-z0-9А-Яа-я]`)
	wordsCount := make(map[string]int)

	arrStrings := strings.Split(input, "\n")

	// firstly split to strings
	for _, line := range arrStrings {
		// then split to words
		words := strings.Split(line, " ")

		for _, word := range words {
			word := strings.ToLower(regex.ReplaceAllString(word, ""))

			if word == "" {
				continue
			}

			wordsCount[word]++
		}
	}

	keys := make([]string, 0, len(wordsCount))

	for k := range wordsCount {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if wordsCount[keys[i]] == wordsCount[keys[j]] {
			return keys[i] < keys[j]
		}

		return wordsCount[keys[i]] > wordsCount[keys[j]]
	})

	if len(keys) > 10 {
		fmt.Printf("%#v\n", keys[:10])
		return keys[:10]
	}

	return keys
}
