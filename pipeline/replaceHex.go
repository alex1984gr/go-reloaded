package pipeline

import (
	"fmt"
	"regexp"
	"strconv"
)

// ReplaceHex εντοπίζει και αντικαθιστά όλους τους hexadecimal αριθμούς (π.χ. 0x1A) με το δεκαδικό τους ισοδύναμο.
func ReplaceHex(lines []string) []string {
	// Κανονική έκφραση που βρίσκει τιμές τύπου 0x1A, 0XFF, 0xabc κλπ.
	re := regexp.MustCompile(`0[xX][0-9a-fA-F]+`)

	var result []string

	for _, line := range lines {
		// Κάνουμε αντικατάσταση για κάθε match
		converted := re.ReplaceAllStringFunc(line, func(hexStr string) string {
			// Μετατρέπουμε το hex σε δεκαδικό
			decValue, err := strconv.ParseInt(hexStr[2:], 16, 64)
			if err != nil {
				// Αν κάτι πάει στραβά, κρατάμε το αρχικό
				return hexStr
			}
			// Επιστρέφουμε τη δεκαδική τιμή ως string
			return fmt.Sprintf("%d", decValue)
		})
		result = append(result, converted)
	}

	return result
}
