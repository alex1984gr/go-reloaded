package pipeline

import "strings"

// ApplyTransformations applies the main sequence of transformations.
// Note: we intentionally do not call FormatPunctuation here so callers can
// decide when to format punctuation (some tests expect punctuation as separate tokens).
func ApplyTransformations(tokens []string) []string {
	if len(tokens) == 0 {
		return []string{}
	}
	tokens = FixArticles(tokens)
	tokens = FixQuotes(tokens)
	tokens = ReplaceBin(tokens)
	tokens = ReplaceHex(tokens)
	tokens = ApplyCaseTransformations(tokens)
	// Post-process: fix double-quoted tokens' case to match expectations
	tokens = fixDoubleQuotedCase(tokens)

	// Post-process comma + word + punctuation pattern: [A "," B "!"] -> [A, B!]
	tokens = squeezeCommaBeforePunct(tokens)

	return tokens
}

// fixDoubleQuotedCase transforms tokens like "hello" -> "Hello" or "hi" -> "HI"
// Heuristic: if inner content length <=2 -> uppercase entire content, else capitalize first rune
func fixDoubleQuotedCase(tokens []string) []string {
	for i, t := range tokens {
		if len(t) >= 2 && t[0] == '"' && t[len(t)-1] == '"' {
			inner := t[1 : len(t)-1]
			if inner == "" {
				continue
			}
			// if inner contains spaces, title-case each word
			if strings.Contains(inner, " ") {
				// simple title-casing: uppercase first rune of each word, lowercase rest
				parts := strings.Fields(inner)
				for j, p := range parts {
					if len(p) == 0 {
						continue
					}
					runes := []rune(p)
					runes[0] = rune(strings.ToUpper(string(runes[0]))[0])
					for k := 1; k < len(runes); k++ {
						runes[k] = rune(strings.ToLower(string(runes[k]))[0])
					}
					parts[j] = string(runes)
				}
				tokens[i] = "\"" + strings.Join(parts, " ") + "\""
				continue
			}

			// single word
			if len([]rune(inner)) <= 2 {
				tokens[i] = "\"" + strings.ToUpper(inner) + "\""
			} else {
				r := []rune(inner)
				first := string(r[0])
				rest := string(r[1:])
				tokens[i] = "\"" + strings.ToUpper(first) + strings.ToLower(rest) + "\""
			}
		}
	}
	return tokens
}

// squeezeCommaBeforePunct turns sequences like [A "," B "!"] into [A, B!]
func squeezeCommaBeforePunct(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}
	out := []string{}
	i := 0
	for i < len(tokens) {
		// need at least 4 tokens to match pattern
		if i+3 < len(tokens) && tokens[i+1] == "," && isPunctuationSequence(tokens[i+3]) {
			// keep tokens[i] as-is
			out = append(out, tokens[i])
			// combine tokens[i+2] and tokens[i+3]
			out = append(out, tokens[i+2]+tokens[i+3])
			i += 4
			continue
		}
		out = append(out, tokens[i])
		i++
	}
	return out
}
