package pipeline

import (
	"strings"
)

// FixQuotes παίρνει ένα slice από tokens και μετατρέπει όλα τα
// words μέσα σε εισαγωγικά σε κεφαλαία.

func FixQuotes(tokens []string) []string {
	inQuotes := false
	result := make([]string, len(tokens))

	for i, tok := range tokens {
		if tok == `"` {
			inQuotes = !inQuotes
			result[i] = tok
			continue
		}
		if inQuotes {
			result[i] = strings.ToUpper(tok)
		} else {
			result[i] = tok
		}
	}
	return result
}
