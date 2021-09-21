package words

import (
	"regexp"
	"sort"
	"strings"
)

const RegexpPattern = `[^\P{P}-]+|( - |- | -)|\s+`

var regExp = regexp.MustCompile(RegexpPattern)

type Word struct {
	Value string
	Count int
}

func AnalyzeText(text string) (words []Word) {
	for word, count := range getCountWords(text) {
		words = append(words, bind(word, count))
	}

	sort.Slice(words, func(i, j int) bool {
		if words[i].Count == words[j].Count {
			return words[i].Value < words[j].Value
		}

		return words[i].Count > words[j].Count
	})

	return words
}

func getWords(text string) []string {
	text = strings.ToLower(regExp.ReplaceAllString(text, " "))

	return strings.Fields(text)
}

func getCountWords(text string) map[string]int {
	count := make(map[string]int)

	for _, word := range getWords(text) {
		count[word]++
	}

	return count
}

func bind(word string, count int) Word {
	return Word{
		Value: word,
		Count: count,
	}
}
