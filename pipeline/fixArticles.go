package pipeline // Defines the package name

import (
	"strings" // Import strings package for basic string operations

	"golang.org/x/text/cases"    // Import for proper Unicode-aware capitalization
	"golang.org/x/text/language" // Import for language settings
)

// FixArticles adjusts "a"/"an" depending on the next word.
// Preserves capitalization if the original article was capitalized.
func FixArticles(words []string) []string {
	if len(words) == 0 { // If the input slice is empty
		return words // Return it unchanged
	}

	result := make([]string, len(words)) // Create a new slice to hold results
	copy(result, words)                  // Copy original words into the result slice

	caser := cases.Title(language.English) // Create a Title-caser for proper capitalization

	for i := 0; i < len(result)-1; i++ { // Loop through all words except the last
		original := result[i]             // Store the original article word
		word := strings.ToLower(original) // Lowercase version for comparison
		next := result[i+1]               // Get the next word

		if word == "a" || word == "an" { // If the current word is an article
			article := "a"                // Default to "a"
			if startsWithVowelOrH(next) { // Check if next word starts with vowel or silent H
				article = "an" // Use "an" if needed
			}

			// Preserve capitalization if original article was capitalized
			if len(original) > 0 && original[0] >= 'A' && original[0] <= 'Z' {
				article = caser.String(article) // Capitalize using x/text/cases
			}

			result[i] = article // Replace the article in the result slice
		}
	}

	return result // Return the modified slice of words
}

// startsWithVowelOrH checks if the word starts with a vowel or a silent H
func startsWithVowelOrH(s string) bool {
	if s == "" { // If the string is empty
		return false // Return false
	}
	s = strings.ToLower(s) // Convert string to lowercase for comparison

	// Handle special silent H words
	if strings.HasPrefix(s, "hour") || strings.HasPrefix(s, "honest") || strings.HasPrefix(s, "honor") {
		return true // Treat as starting with vowel
	}

	first := s[0]                                     // Get the first character
	return strings.ContainsRune("aeiou", rune(first)) // Return true if it's a vowel
}
