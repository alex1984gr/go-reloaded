package pipeline // This file belongs to the 'pipeline' package

import (
	"os" // Import the 'os' package to interact with the operating system (for reading files)
)

// ReadInput reads the entire content of a file and returns it as a string.
// If an error occurs (e.g., file not found), it returns an empty string and the error.
func ReadInput(path string) (string, error) { // Function named ReadInput takes a file path and returns a string and an error
	data, err := os.ReadFile(path) // Read the entire file content from the given path; returns byte slice and error
	if err != nil {                // Check if there was an error during reading
		return "", err // If error exists, return an empty string and the error
	}
	return string(data), nil // Convert byte slice to string and return it with no error (nil)
}
