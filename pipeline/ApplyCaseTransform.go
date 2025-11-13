package pipeline // This belongs to the "pipeline" package

import (
	"strconv"
	"strings"
	"unicode"
)

// ApplyCaseTransformations applies all uppercase, lowercase, and capitalize transformations
// based on markers: (up), (low), (cap) or (up, N), (low, N), (cap, N)
func ApplyCaseTransformations(tokens []string) []string {
	var result []string // Resulting token slice

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Check if token starts with "(" and ends with ")" indicating a marker
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {
			// Remove parentheses
			content := token[1 : len(token)-1]

			// Split by comma to check if there's a number argument
			parts := strings.Split(content, ",")
			action := strings.TrimSpace(parts[0]) // "up", "low", or "cap"
			count := 1                            // Default: 1 word

			if len(parts) == 2 { // If there's a number
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}

			// Apply transformation to the last 'count' words in result
			switch strings.ToLower(action) {
			case "up":
				applyUppercase(result, count)
			case "low":
				applyLowercase(result, count)
			case "cap":
				applyCapitalize(result, count)
			default:
				// Unknown marker, just ignore
			}

			// Skip adding the marker itself
			continue
		}

		// Normal token, add to result
		result = append(result, token)
	}

	return result
}

// Helper: apply uppercase to last 'count' tokens in slice
func applyUppercase(tokens []string, count int) {
	start := max(0, len(tokens)-count)
	for i := start; i < len(tokens); i++ {
		tokens[i] = strings.ToUpper(tokens[i])
	}
}

// Helper: apply lowercase to last 'count' tokens in slice
func applyLowercase(tokens []string, count int) {
	start := max(0, len(tokens)-count)
	for i := start; i < len(tokens); i++ {
		tokens[i] = strings.ToLower(tokens[i])
	}
}

// Helper: capitalize last 'count' tokens in slice
func applyCapitalize(tokens []string, count int) {
	start := max(0, len(tokens)-count)
	for i := start; i < len(tokens); i++ {
		tokens[i] = capitalizeWord(tokens[i])
	}
}

// Capitalize the first letter of a word, keep the rest lowercase
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// Utility: max function to avoid negative indices
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
