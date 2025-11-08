package pipeline

func ApplyTransformations(input []string) []string {
	if len(input) == 0 {
		return input
	}

	result := input

	result = FixArticles(result)
	result = FixQuotes(result)
	result = ReplaceBin(result)
	result = ReplaceHex(result)
	result = ApplyCaseTransform(result, "capitalize")
	result = FormatPunctuation(result)
	return result
}
