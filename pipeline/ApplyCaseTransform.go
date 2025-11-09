package pipeline

import (
	"strings"
	"unicode"
)

// ApplyCaseTransform εφαρμόζει σωστά κεφαλαία και μικρά ανάλογα με quotes
func ApplyCaseTransform(tokens []string) []string {
	result := make([]string, len(tokens))
	inQuotes := false

	for i, token := range tokens {
		switch token {
		case `"`:
			inQuotes = !inQuotes
			result[i] = token
		default:
			if inQuotes {
				result[i] = toUpper(token)
			} else {
				result[i] = capitalizeWord(token)
			}
		}
	}
	return result
}

// capitalizeWord: κεφαλαιοποιεί μόνο την πρώτη γράμμα της λέξης
func capitalizeWord(token string) string {
	if token == "" {
		return ""
	}

	runes := []rune(token)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// toUpper: μετατρέπει όλα τα γράμματα σε κεφαλαία
func toUpper(token string) string {
	return strings.Map(func(r rune) rune {
		return unicode.ToUpper(r)
	}, token)
}
