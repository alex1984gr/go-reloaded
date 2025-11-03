package pipeline

func Tokenize(input []rune) []string {
	tokens := []string{} // slice για να αποθηκεύσει τα tokens
	current := ""        // προσωρινή μεταβλητή για την τρέχουσα λέξη

	for _, ch := range input { // loop σε κάθε χαρακτήρα του input
		if ch == '<' || ch == '>' || ch == '"' || ch == ' ' { // αν είναι ειδικός χαρακτήρας
			if current != "" { // αν υπάρχει κάτι στο current
				tokens = append(tokens, current) // προσθέτει την τρέχουσα λέξη στα tokens
				current = ""                     // καθαρίζει το current
			}
			if ch != ' ' { // αν δεν είναι κενό
				tokens = append(tokens, string(ch)) // προσθέτει τον ειδικό χαρακτήρα στα tokens
			}
		} else {
			current += string(ch) // αλλιώς προσθέτει τον χαρακτήρα στο current
		}
	}

	if current != "" { // αν υπάρχει υπόλοιπο στο current
		tokens = append(tokens, current) // το προσθέτει στα tokens
	}

	return tokens // επιστρέφει όλα τα tokens
}
