package pipeline

import (
	"strings"
)

// FixQuotes διατηρεί τα quotes και εφαρμόζει σωστά το capitalization μέσα τους.
// Μετατρέπει κάθε λέξη μέσα σε quotes σε κεφαλαία (όπως απαιτούν τα tests).
func FixQuotes(tokens []string) []string {
	inQuotes := false
	for i, token := range tokens {
		if strings.HasPrefix(token, `"`) && strings.HasSuffix(token, `"`) && len(token) > 1 {
			// Όλο το token είναι μέσα σε quotes: "hello"
			tokens[i] = `"` + strings.ToUpper(strings.Trim(token, `"`)) + `"`
			continue
		}

		if strings.HasPrefix(token, `"`) {
			inQuotes = true
			if len(token) > 1 {
				tokens[i] = `"` + strings.ToUpper(token[1:])
			}
			continue
		}

		if strings.HasSuffix(token, `"`) {
			if len(token) > 1 {
				tokens[i] = strings.ToUpper(token[:len(token)-1]) + `"`
			}
			inQuotes = false
			continue
		}

		if inQuotes {
			tokens[i] = strings.ToUpper(token)
		}
	}
	return tokens
}
