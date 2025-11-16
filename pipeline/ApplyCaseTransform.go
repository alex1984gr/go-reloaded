package pipeline // Declares that this file belongs to the "pipeline" package

import (
	"strconv" // Used for converting string numbers to integers
	"strings" // Provides utilities for splitting and formatting text
	"unicode" // Provides rune-level character transformations (upper/lower)
)

// ApplyCaseTransformations processes the token list and applies case-transform rules.
// Rules are encoded in tokens like: (up), (low), (cap), (up, N), (low, N), (cap, N)
func ApplyCaseTransformations(tokens []string) []string {
	var result []string // Accumulates all final tokens after processing

	for i := 0; i < len(tokens); i++ { // Iterate through each token in the input
		token := tokens[i] // Current token under inspection

		// Check whether the token is a transformation marker, e.g. "(up)" or "(cap, 3)"
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {

			// Extract the inner content without the parentheses
			content := token[1 : len(token)-1]

			// Some markers include a count after a comma (e.g. "up, 3")
			parts := strings.Split(content, ",")
			action := strings.TrimSpace(parts[0]) // The instruction keyword: up/low/cap
			count := 1                            // Default number of affected tokens

			// If a count is provided, parse it
			if len(parts) == 2 {
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}

			// Apply the correct transformation to the previously collected tokens
			// depending on the requested action
			switch strings.ToLower(action) {
			case "up":
				applyUppercase(result, count)
			case "low":
				applyLowercase(result, count)
			case "cap":
				applyCapitalize(result, count)
			default:
				// Unknown instruction: ignore silently
			}

			// Do NOT append the marker itself to output
			continue
		}

		// Normal token — push it to the result slice
		result = append(result, token)
	}

	return result // Return the fully processed token sequence
}

// applyUppercase converts the last <count> tokens of the slice to uppercase.
func applyUppercase(tokens []string, count int) {
	start := max(0, len(tokens)-count) // Determine the safe starting index
	for i := start; i < len(tokens); i++ {
		tokens[i] = strings.ToUpper(tokens[i])
	}
}

// applyLowercase converts the last <count> tokens of the slice to lowercase.
func applyLowercase(tokens []string, count int) {
	start := max(0, len(tokens)-count)
	for i := start; i < len(tokens); i++ {
		tokens[i] = strings.ToLower(tokens[i])
	}
}

// applyCapitalize capitalizes the last <count> tokens.
// Capitalizing means: first letter uppercase, remaining letters lowercase.
func applyCapitalize(tokens []string, count int) {
	start := max(0, len(tokens)-count)
	for i := start; i < len(tokens); i++ {
		tokens[i] = capitalizeWord(tokens[i])
	}
}

// capitalizeWord transforms a single token so that:
// - the first character is uppercase
// - the remaining characters are lowercase
func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word // No operation on empty tokens
	}

	runes := []rune(word) // Convert string to rune slice for safe Unicode operations

	runes[0] = unicode.ToUpper(runes[0]) // Uppercase first character
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i]) // Lowercase the rest
	}

	return string(runes) // Convert back to string
}

// max returns the larger of two integers.
// Used to avoid negative slice indexing.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
