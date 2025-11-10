package main

import (
	"fmt"
	"log"
	"os"

	"go-reloaded/pipeline"
)

func main() {
	// ✅ Έλεγχος ορισμάτων
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// 🧩 Βήμα 1: Ανάγνωση του αρχείου εισόδου
	text, err := pipeline.ReadInput(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// 🧩 Βήμα 2: Tokenization (χωρίζει σε λέξεις, σημεία στίξης κλπ.)
	tokens := pipeline.Tokenize([]rune(text))

	// 🧩 Βήμα 3: Εφαρμογή όλων των μετασχηματισμών
	tokens = pipeline.ReplaceHex(tokens)
	tokens = pipeline.ReplaceBin(tokens)
	tokens = pipeline.ApplyCaseTransform(tokens)
	tokens = pipeline.FormatPunctuation(tokens)
	tokens = pipeline.FixQuotes(tokens)
	tokens = pipeline.FixArticles(tokens)

	// 🧩 Βήμα 4: Επανασύνθεση του τελικού κειμένου
	result := pipeline.JoinTokens(tokens)

	// 🧩 Βήμα 5: Εγγραφή στο αρχείο εξόδου
	err = pipeline.WriteOutput(outputFile, []string{result})
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Println("✅ File processed successfully:", outputFile)
}
