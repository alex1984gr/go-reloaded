package pipeline

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// FixArticles διορθώνει τα άρθρα "a"/"an" ανάλογα με το αν η επόμενη λέξη ξεκινά από φωνήεν ή "h"
func FixArticles(words []string) []string {
	if len(words) == 0 {
		return words
	}

	result := make([]string, len(words))
	copy(result, words)

	caser := cases.Title(language.English)

	for i := 0; i < len(result)-1; i++ {
		word := strings.ToLower(result[i])
		next := strings.ToLower(result[i+1])

		if word == "a" || word == "an" {
			if startsWithVowelOrH(next) {
				result[i] = "an"
			} else {
				result[i] = "a"
			}
		}

		// Αν η αρχική λέξη ήταν κεφαλαία, διατήρησε το κεφαλαίο άρθρο
		if caser.String(result[i]) == "A" {
			result[i] = "A"
		} else if caser.String(result[i]) == "An" {
			result[i] = "An"
		}
	}

	return result
}

// startsWithVowelOrH ελέγχει αν η λέξη ξεκινάει από φωνήεν ή 'h'
func startsWithVowelOrH(s string) bool {
	if s == "" {
		return false
	}
	first := s[0]
	return strings.ContainsRune("aeiouh", rune(first))
}
