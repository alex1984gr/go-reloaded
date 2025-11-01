package tests

import (
	"go-reloaded/pipeline"
	"os"
	"testing"
)

func TestReadInput(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "input_*.txt") // Δημιουργούμε προσωρινό αρχείο για το test
	if err != nil {
		t.Fatalf("❌ Αποτυχία δημιουργίας προσωρινού αρχείου: %v", err) // Διαγράφουμε το αρχείο μετά το τέλος του test
	}
	defer os.Remove(tmpFile.Name())

	// Γράφουμε μέσα του ένα απλό περιεχόμενο για δοκιμή
	content := "Hello, Go Reloaded!"
	tmpFile.WriteString(content)
	tmpFile.Close()

	// Καλούμε τη συνάρτηση readInput από το pipeline
	result, err := pipeline.ReadInput(tmpFile.Name())

	// Ελέγχουμε αν επιστράφηκε error (δεν θα έπρεπε)
	if err != nil {
		t.Fatalf("❌ Αναπάντεχο error: %v", err)
	}

	// Ελέγχουμε αν το αποτέλεσμα είναι το ίδιο με το αρχικό περιεχόμενο
	if result != content {
		t.Errorf("❌ Αναμενόμενο '%s', έλαβα '%s'", content, result)
	}

	// Αν φτάσει εδώ, όλα πήγαν καλά
	t.Logf("✅ Το readInput διάβασε σωστά το αρχείο: %s", tmpFile.Name())
}

// TestReadInput_FileNotFound ελέγχει τη συμπεριφορά όταν το αρχείο δεν υπάρχει
func TestReadInput_FileNotFound(t *testing.T) {
	// Καλούμε τη συνάρτηση με path που δεν υπάρχει
	_, err := pipeline.ReadInput("non_existing_file.txt")

	// Αν δεν επιστρέψει error, αποτυγχάνει το test
	if err == nil {
		t.Errorf("❌ Αναμενόταν σφάλμα για ανύπαρκτο αρχείο, αλλά δεν ελήφθη.")
	} else {
		t.Logf("✅ Ο χειρισμός ανύπαρκτου αρχείου λειτουργεί σωστά: %v", err)
	}
}
