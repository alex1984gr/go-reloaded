package pipeline

import "strings"

// FormatPunctuation removes spaces before punctuation and attaches punctuation to previous word.
func FormatPunctuation(tokens []string) []string {
	var out []string
	punct := ",.!?;:"

	for _, t := range tokens {
		if len(t) == 1 && strings.Contains(punct, t) {
			if len(out) > 0 {
				out[len(out)-1] += t
			} else {
				out = append(out, t)
			}
		} else {
			out = append(out, t)
		}
	}
	return out
}
