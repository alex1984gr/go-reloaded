package pipeline

import (
	"strings"
)

// FormatPunctuation formats punctuation according to the rules:
// 1. Basic punctuation (.,!?) sticks to previous word and has space after next word.
// 2. Colon (:) and semicolon (;) remain separate tokens.
// 3. Multiple punctuation marks like !! or ?? are merged with previous word.
// 4. Ellipsis ... and combined punctuation like !? are merged with previous word.
// 5. Single quotes '...' wrap words or phrases without extra spaces.
func FormatPunctuation(tokens []string) []string {
	var result []string      // Output tokens
	var quoteBuffer []string // Buffer to collect words inside quotes
	inQuote := false         // Flag to indicate if we're inside quotes

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// Handle start or end of single quotes
		if token == "'" {
			if !inQuote {
				// Start of quote: begin collecting words inside quotes
				inQuote = true
				quoteBuffer = []string{}
			} else {
				// End of quote: flush buffer as single quoted token
				inQuote = false
				result = append(result, "'"+strings.Join(quoteBuffer, " ")+"'")
			}
			continue // Skip the quote token itself
		}

		if inQuote {
			// Collect words inside quotes
			quoteBuffer = append(quoteBuffer, token)
			continue
		}

		// Handle ellipsis ... or combined punctuation !?
		if token == "..." || token == "!?" {
			if len(result) > 0 {
				result[len(result)-1] += token // Attach to previous word
			} else {
				result = append(result, token)
			}
			continue
		}

		// Handle basic punctuation that should attach to previous word
		if isBasicPunctuation(token) {
			if len(result) > 0 && !isColonOrSemicolon(token) {
				// Attach punctuation to previous word
				result[len(result)-1] += token
			} else {
				// Colon or semicolon remain separate
				result = append(result, token)
			}
			continue
		}

		// Merge multiple punctuation tokens like !! or ??
		for i+1 < len(tokens) && isRepeatedPunctuation(tokens[i+1]) {
			token += tokens[i+1] // Merge with current token
			i++                  // Skip the next token since merged
		}

		// Normal token: append to result
		result = append(result, token)
	}

	return result
}

// Check if token is basic punctuation: . , ! ?
func isBasicPunctuation(token string) bool {
	basic := []string{".", ",", "!", "?"}
	for _, p := range basic {
		if token == p {
			return true
		}
	}
	return false
}

// Colon and semicolon stay separate
func isColonOrSemicolon(token string) bool {
	return token == ":" || token == ";"
}

// Check for repeated punctuation like !! or ??
func isRepeatedPunctuation(token string) bool {
	if len(token) < 2 {
		return false
	}
	first := token[0]
	for i := 1; i < len(token); i++ {
		if token[i] != first {
			return false
		}
	}
	return strings.Contains(".!?,", string(first)) // Only punctuation allowed
}
