package pipeline

func FormatPunctuation(tokens []string) []string { // Function that formats punctuation tokens correctly
	if len(tokens) == 0 { // If there are no tokens, just return them as-is
		return tokens
	}

	var result []string            // Slice to store the processed tokens
	for _, token := range tokens { // Loop through each token in the input
		if token == " " { // If the token is just a space
			continue // Skip it
		}
		if isPunctuation(token) && len(result) > 0 { // If the token is punctuation and there’s something before it
			result[len(result)-1] += token // Attach the punctuation directly to the previous word (no space)
		} else { // Otherwise, it’s a normal word or first token
			result = append(result, token) // Add it as a new token to the result
		}
	}

	return result // Return the cleaned-up list of tokens
}

func isPunctuation(token string) bool { // Helper function to check if a token is punctuation
	switch token { // Compare the token against known punctuation marks
	case ".", ",", "!", "?", ";", ":": // Common punctuation symbols
		return true // If matched, return true
	default:
		return false // Otherwise, return false
	}
}
