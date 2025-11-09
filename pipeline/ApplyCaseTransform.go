package pipeline

import (
	"strings"
	"unicode"
)

func ApplyCaseTransform(tokens []string) []string {
	inQuotes := false
	result := make([]string, len(tokens))

	for i, tok := range tokens {
		// Αν η λέξη είναι quote
		if tok == `"` {
			inQuotes = !inQuotes
			result[i] = tok
			continue
		}

		// Αν είμαστε μέσα σε quotes, κάνουμε όλα τα γράμματα uppercase
		if inQuotes {
			result[i] = strings.ToUpper(tok)
		} else {
			// Κανονικό capitalize: πρώτο γράμμα κεφαλαίο, υπόλοιπα μικρά
			if len(tok) > 0 {
				runes := []rune(tok)
				runes[0] = unicode.ToUpper(runes[0])
				for j := 1; j < len(runes); j++ {
					runes[j] = unicode.ToLower(runes[j])
				}
				result[i] = string(runes)
			} else {
				result[i] = tok
			}
		}
	}

	return result
}
