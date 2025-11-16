package pipeline

import (
	"strings"
)

// FormatPunctuation applies punctuation formatting rules:
// - groups of punctuation (e.g. "...", "!!", "?!") attach to previous token
// - single punctuation marks attach to previous token
// - handle single quotes via FixQuotes
// - trim double spaces
func FormatPunctuation(tokens []string) []string {
	// Attach any punctuation-only tokens (made of .,!?:;) to the previous token
	if len(tokens) == 0 {
		return tokens
	}

	var out []string
	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if isPunctuationSequence(t) && len(out) > 0 {
			// attach to previous token
			out[len(out)-1] = out[len(out)-1] + t
		} else {
			out = append(out, t)
		}
	}

	// Now fix quotes and clean spaces
	out = FixQuotes(out)
	for i := 0; i < len(out); i++ {
		// collapse multiple spaces and trim
		s := out[i]
		for strings.Contains(s, "  ") {
			s = strings.ReplaceAll(s, "  ", " ")
		}
		out[i] = strings.TrimSpace(s)
	}

	return out
}

// isPunctuationSequence returns true if the token consists only of the
// punctuation characters we care about: . , ! ? : ; (one or more times)
func isPunctuationSequence(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		switch r {
		case '.', ',', '!', '?', ':', ';':
			// allowed
		default:
			return false
		}
	}
	return true
}
