package hw03frequencyanalysis

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// []string
func Top10(input string) map[string]int {
	regex, _ := regexp.Compile(`[^A-Za-z0-9А-Яа-я]`)
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

			wordsCount[word] = wordsCount[word] + 1

		}
	}

	keys := make([]string, 0, len(wordsCount))

	for k, _ := range wordsCount {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if wordsCount[keys[i]] == wordsCount[keys[j]] {
			return keys[i] < keys[j]
		}

		return wordsCount[keys[i]] > wordsCount[keys[j]]
	})

	// todo return first 10

	fmt.Printf("%#v\n", keys)

	return wordsCount
}
