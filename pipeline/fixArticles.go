package pipeline // Defines the package name

import (
	"strings" // For string operations

	"golang.org/x/text/cases"    // For proper Unicode-aware capitalization
	"golang.org/x/text/language" // For language settings
)

// FixArticles converts "a" to "an" if the next word starts with a vowel or 'h'.
// Preserves capitalization of the original article.
func FixArticles(words []string) []string {
	if len(words) == 0 { // Return empty slice if input is empty
		return words
	}

	result := make([]string, len(words)) // Prepare output slice
	copy(result, words)                  // Copy input to result

	caser := cases.Title(language.English) // Title-caser for capitalized articles

	for i := 0; i < len(result)-1; i++ { // Loop through all words except last
		original := result[i]             // Store the current word
		word := strings.ToLower(original) // Lowercase for comparison
		next := result[i+1]               // Get next word

		// Only care about "a" (or capitalized "A")
		if word == "a" {
			article := "a" // Default

			// If next word starts with vowel or 'h', use "an"
			if startsWithVowelOrH(next) {
				article = "an"
			}

			// Preserve capitalization
			if len(original) > 0 && original[0] >= 'A' && original[0] <= 'Z' {
				article = caser.String(article)
			}

			result[i] = article // Replace in result
		}
	}

	return result
}

// startsWithVowelOrH checks if a word starts with a vowel or 'h'
func startsWithVowelOrH(s string) bool {
	if s == "" {
		return false
	}
	first := strings.ToLower(string(s[0])) // Lowercase first character
	return strings.Contains("aeiouh", first)
}
