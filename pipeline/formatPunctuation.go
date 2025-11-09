package pipeline

import (
	"strings"
)

func FormatPunctuation(tokens []string) []string {
	var result []string
	punctuation := ",.!?;:"
	for _, tok := range tokens {
		if len(tok) == 1 && strings.Contains(punctuation, tok) && len(result) > 0 {
			// Προσάρμοσε το punctuation στην προηγούμενη λέξη χωρίς να αλλάξεις capitalization
			result[len(result)-1] += tok
		} else if tok != " " {
			result = append(result, tok)
		}
	}
	return result
}
