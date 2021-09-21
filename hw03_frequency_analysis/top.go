package hw03frequencyanalysis

import (
	"github.com/i4erkasov/otus_hw/hw03_frequency_analysis/words"
)

func Top10(text string) []string {
	return getTop(text, 10)
}

func getTop(text string, limit int) (top []string) {
	wordsWithCount := words.AnalyzeText(text)

	if len(wordsWithCount) < limit {
		limit = len(wordsWithCount)
	}

	for i := 0; i < limit; i++ {
		top = append(top, wordsWithCount[i].Value)
	}

	return top
}
