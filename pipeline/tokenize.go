package pipeline

func Tokenize(input []rune) []string {
	tokens := []string{} // slice to store the tokens
	current := ""        // temporary variable to hold the current word

	for _, ch := range input { // loop through each character in the input
		if ch == '<' || ch == '>' || ch == '"' || ch == ' ' { // if it’s a special character
			if current != "" { // if there’s something stored in 'current'
				tokens = append(tokens, current) // add the current word to the tokens
				current = ""                     // clear 'current'
			}
			if ch != ' ' { // if it’s not a space
				tokens = append(tokens, string(ch)) // add the special character as a token
			}
		} else {
			current += string(ch) // otherwise, add the character to 'current'
		}
	}

	if current != "" { // if there’s anything left in 'current'
		tokens = append(tokens, current) // add it to the tokens
	}

	return tokens // return all tokens
}
