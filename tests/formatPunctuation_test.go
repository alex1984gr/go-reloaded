package tests // Define the package for test files

import (
	"fmt"                  // Used for printing debug information
	"go-reloaded/pipeline" // Import the pipeline functions to test
	"testing"              // Go's testing framework
)

// TestFormatPunctuation runs all punctuation formatting tests
func TestFormatPunctuation(t *testing.T) {

	// debug is a helper function to print verbose debug logs
	debug := func(name string, input, output, expected []string) {
		fmt.Println()                                     // Empty line for readability
		fmt.Println("[DEBUG] Running test:", name)        // Print the test name
		fmt.Println("[DEBUG] Input tokens:   ", input)    // Show the raw input tokens
		fmt.Println("[DEBUG] Output tokens:  ", output)   // Tokens after calling FormatPunctuation
		fmt.Println("[DEBUG] Expected tokens:", expected) // Expected correct result
		fmt.Println()                                     // Extra spacing for clarity
	}

	// ---- TEST 1: Basic punctuation spacing ----
	t.Run("Basic punctuation spacing", func(t *testing.T) {
		name := "Basic punctuation spacing" // Name for debug printing

		// Input tokens as the tokenizer produces them
		input := []string{"I", "was", "sitting", "over", "there", ",", "and", "then", "BAMM", "!!"}

		// Expected transformation according to the rules
		expected := []string{"I", "was", "sitting", "over", "there,", "and", "then", "BAMM!!"}

		// Process the tokens with FormatPunctuation
		output := pipeline.FormatPunctuation(input)

		// Print debug logs
		debug(name, input, output, expected)

		// Compare output with expected using a helper function
		if !equalSlices(output, expected) {
			// If mismatch, mark test as failed with a detailed message
			t.Errorf("❌ Test failed: %s\nExpected: %v\nGot:      %v",
				name, expected, output)
		} else {
			fmt.Println("✅ Test passed:", name) // On success, print a nice check mark
		}
	})

	// ---- TEST 2: Handling "..." ellipsis ----
	t.Run("Ellipsis spacing", func(t *testing.T) {
		name := "Ellipsis spacing"

		// Input where "..." should stick to the previous word
		input := []string{"I", "was", "thinking", "...", "You", "were", "right"}

		// Expected: "thinking..." becomes a single token attached to the previous
		expected := []string{"I", "was", "thinking...", "You", "were", "right"}

		output := pipeline.FormatPunctuation(input)

		debug(name, input, output, expected)

		if !equalSlices(output, expected) {
			t.Errorf("❌ Test failed: %s\nExpected: %v\nGot:      %v",
				name, expected, output)
		} else {
			fmt.Println("✅ Test passed:", name)
		}
	})

	// ---- TEST 3: Combined punctuation like "!?" ----
	t.Run("Combined punctuation spacing", func(t *testing.T) {
		name := "Combined punctuation spacing"

		// Combined punctuation should stay together and attach to previous word
		input := []string{"Wait", "!?"}

		expected := []string{"Wait!?"}

		output := pipeline.FormatPunctuation(input)

		debug(name, input, output, expected)

		if !equalSlices(output, expected) {
			t.Errorf("❌ Test failed: %s\nExpected: %v\nGot:      %v",
				name, expected, output)
		} else {
			fmt.Println("✅ Test passed:", name)
		}
	})

	// ---- TEST 4: A single-word inside quotes ----
	t.Run("Single word quote", func(t *testing.T) {
		name := "Single word quote"

		// Input produced by the tokenizer: quotes become separate tokens
		input := []string{"I", "am", "exactly", "how", "they", "describe", "me", ":", "'", "awesome", "'"}

		// Expected: "'awesome'" should become a single compact token
		expected := []string{"I", "am", "exactly", "how", "they", "describe", "me", ":", "'awesome'"}

		output := pipeline.FormatPunctuation(input)

		debug(name, input, output, expected)

		// Check if the output matches expected behavior
		if !equalSlices(output, expected) {
			t.Errorf("❌ Test failed: %s\nExpected: %v\nGot:      %v",
				name, expected, output)
		} else {
			fmt.Println("✅ Test passed:", name)
		}
	})

	// ---- TEST 5: Multi-word inside quotes ----
	t.Run("Multiple word quote", func(t *testing.T) {
		name := "Multiple word quote"

		// Input: multiple words between quotes
		input := []string{"As", "Elton", "John", "said", ":", "'", "I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world", "'"}

		// Expected: quotes attach to the first & last word
		expected := []string{"As", "Elton", "John", "said", ":", "'I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world'"}

		output := pipeline.FormatPunctuation(input)

		debug(name, input, output, expected)

		if !equalSlices(output, expected) {
			t.Errorf("❌ Test failed: %s\nExpected: %v\nGot:      %v",
				name, expected, output)
		} else {
			fmt.Println("✅ Test passed:", name)
		}
	})

}

// equalSlices compares two string slices element-by-element
func equalSlices(a, b []string) bool {
	if len(a) != len(b) { // If lengths differ, slices are not equal
		return false
	}
	for i := range a { // Iterate through slice
		if a[i] != b[i] { // Compare each element
			return false // Mismatch
		}
	}
	return true // All elements match
}
