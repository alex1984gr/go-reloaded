package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"go-reloaded/pipeline"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Διαβάζουμε όλο το input
	text, err := pipeline.ReadInput(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Σπάμε σε γραμμές
	lines := strings.Split(text, "\n")

	var results []string

	for _, line := range lines {
		if line == "" {
			results = append(results, "")
			continue
		}

		tokens := pipeline.Tokenize([]rune(line))

		tokens = pipeline.ReplaceHex(tokens)
		tokens = pipeline.ReplaceBin(tokens)
		tokens = pipeline.ApplyCaseTransformations(tokens)
		tokens = pipeline.FormatPunctuation(tokens)
		tokens = pipeline.FixQuotes(tokens)
		tokens = pipeline.FixArticles(tokens)

		result := pipeline.JoinTokens(tokens)
		results = append(results, result)
	}

	// Γράφουμε ΚΑΘΕ γραμμή στο output
	err = pipeline.WriteOutput(outputFile, results)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Println("✅ File processed successfully:", outputFile)
}
