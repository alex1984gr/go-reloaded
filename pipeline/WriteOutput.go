package pipeline

import (
	"os"
	"strings"
)

// Ενώνει τα tokens σε ένα ενιαίο string.
// Τα σημεία στίξης έχουν ήδη διαμορφωθεί σωστά από FormatPunctuation,
// οπότε το Join γίνεται με space.
func JoinTokens(tokens []string) string {
	return strings.Join(tokens, " ")
}

// Γράφει το output στο αρχείο.
// Κάθε στοιχείο του input slice γίνεται δική του γραμμή.
// Το αρχείο τελειώνει με newline όπως απαιτεί η άσκηση.
func WriteOutput(filename string, input []string) error {
	output := strings.Join(input, "\n") + "\n"
	return os.WriteFile(filename, []byte(output), 0o644)
}
