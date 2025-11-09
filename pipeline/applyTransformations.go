package pipeline

func ApplyTransformations(tokens []string) []string {
	if len(tokens) == 0 {
		return tokens
	}

	// 1️⃣ Fix articles (a → an)
	tokens = FixArticles(tokens)

	// 2️⃣ Fix quotes (words inside quotes to uppercase)
	tokens = FixQuotes(tokens)

	// 3️⃣ Replace binary numbers with decimals
	tokens = ReplaceBin(tokens)

	// 4️⃣ Replace hex numbers with decimals
	tokens = ReplaceHex(tokens)

	// 5️⃣ Apply case transformation (capitalize words outside quotes, keep inside quotes)
	tokens = ApplyCaseTransform(tokens)

	// 6️⃣ Format punctuation (remove spaces before .,!? etc.)
	tokens = FormatPunctuation(tokens)

	return tokens
}
