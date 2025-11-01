package pipeline

import (
	"bufio"
	"os"
)

func ReadInput(s string) []rune {
	scanner := bufio.NewScanner(os.Stdin) // δημιουργεί ένα σκάνερ που διαβάζει γραμμές από το τερματικό
	text := ""                            // δημιουργεί μεταβλητή text τύπου string για αποθήκευση κειμένου
	for scanner.Scan() {                  // loop που τρέχει όσο υπάρχουν γραμμές για να διαβάσει ο scanner
		text += scanner.Text() + " " // προσθέτει τη γραμμή που διάβασε το scanner στο string text
	}
	return []rune(text) // μετατρέπει όλο το string σε slice από runes
}
