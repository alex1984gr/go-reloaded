package pipeline

// Tokenize splits the input runes into tokens used by the pipeline.
// Rules:
// - whitespace splits tokens
// - punctuation sequences (.,!?:;) are returned as single tokens (e.g. "...", "!!")
// - single and double quotes, angle brackets returned as single tokens
// - markers like "(up)" or "(hex)" are kept as a single token (no internal split)
func Tokenize(input []rune) []string {
	var tokens []string
	n := len(input)
	i := 0

	isPunct := func(r rune) bool {
		switch r {
		case '.', ',', '!', '?', ':', ';':
			return true
		}
		return false
	}

	for i < n {
		r := input[i]

		// whitespace -> skip and flush
		if r == ' ' || r == '\n' || r == '\t' || r == '\r' {
			i++
			continue
		}

			// markers like (up) or ( low, 3 ): capture everything until the next ')'
			if r == '(' {
				j := i + 1
				valid := false
				for j < n {
					if input[j] == ')' {
						valid = true
						break
					}
					j++
				}
				if valid {
					// capture whole parentheses, preserving inner spaces (we'll normalize later)
					tokens = append(tokens, string(input[i:j+1]))
					i = j + 1
					continue
				}
			}

		// punctuation sequences
		if isPunct(r) {
			j := i + 1
			for j < n && isPunct(input[j]) {
				j++
			}
			tokens = append(tokens, string(input[i:j]))
			i = j
			continue
		}

		// single/double quote or angle bracket as single token
		if r == '\'' || r == '"' || r == '<' || r == '>' {
			tokens = append(tokens, string(r))
			i++
			continue
		}

		// otherwise collect a word until next separator
		j := i
		for j < n {
			rr := input[j]
			if rr == ' ' || rr == '\n' || rr == '\t' || rr == '\r' {
				break
			}
			if rr == '(' || rr == ')' || rr == '\'' || rr == '"' || rr == '<' || rr == '>' || isPunct(rr) {
				break
			}
			j++
		}
		if j > i {
			tokens = append(tokens, string(input[i:j]))
			i = j
		} else {
			// fallback single rune token
			tokens = append(tokens, string(input[i]))
			i++
		}
	}

	return tokens
}
