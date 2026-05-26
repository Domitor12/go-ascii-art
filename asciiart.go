// Package declaration - every Go file must start with this
// "main" tells Go this is an executable program, not a library
// This is required for Go to know where program execution begins
package main

// Import section - bringing in tools from Go's standard library
// These packages provide essential functionality our program needs
import (
	"fmt"     // Format package - provides printing functions (fmt.Println, fmt.Printf)
	"os"      // Operating System package - for reading files and accessing command line arguments
	"strings" // Strings package - provides text manipulation functions (Split, ReplaceAll, etc.)
)

// Each ASCII character is 8 lines tall
// This is a constant because ALL banner files (standard, shadow, thinkertoy) use 8 lines per character
// Using a named constant makes the code more readable and maintainable
// If the format ever changes, we only need to update this one value
const charHeight = 8

// ASCII table starts from space ' ' (code 32)
// Printable ASCII characters range from 32 (space) to 126 (tilde)
// We store space at index 0 in our banner parsing, so we subtract this value
// This constant makes the calculation clear: index = asciiCode - asciiStart
const asciiStart = 32

// Each character block = 8 lines + 1 empty separator
// In the banner file format, each character occupies 9 lines:
// - 8 lines for the actual ASCII art
// - 1 empty line as a separator between characters
// This makes it easier to parse individual characters from the banner file
const blockSize = 9

// Function to load banner file into memory
// Takes a filename (e.g., "standard.txt") and returns the file content as lines
// Returns two values: the lines slice and any error that occurred
// This is a helper function that encapsulates file reading logic
func loadBanner(filename string) ([]string, error) {
	// os.ReadFile reads the entire file into memory
	// Returns the file data as bytes and any error (file not found, permissions, etc.)
	data, err := os.ReadFile(filename)

	// Check if there was an error reading the file
	// In Go, we always check errors immediately after they occur
	if err != nil {
		// Return nil for the data and the error to the caller
		// The caller will handle the error appropriately
		return nil, err
	}

	// Convert the byte slice to a string and split by newline characters
	// strings.Split divides the text wherever it finds "\n"
	// This creates a slice where each element is one line from the file
	// Return the lines slice and nil (no error)
	return strings.Split(string(data), "\n"), nil
}

// Function to get ASCII-art lines for a character
// Takes a character (rune) and the banner lines slice
// Returns the 8 lines that form the ASCII art for that character
// This function encapsulates the logic of locating a character in the banner file
func getCharArt(char rune, bannerLines []string) []string {
	// Calculate the index where this character's art starts in the banner file
	// Formula: (ASCII_code - starting_ASCII) * block_size
	// Example: 'A' (65) -> (65 - 32) * 9 = 33 * 9 = 297
	// This means 'A' starts at line 297 in the banner file
	index := (int(char) - asciiStart) * blockSize

	// Return a slice of 8 lines starting at the calculated index
	// This extracts exactly the 8 lines that form the character's ASCII art
	// bannerLines[index : index+charHeight] means "from index to index+7"
	return bannerLines[index : index+charHeight]
}

// Function to print one word in ASCII-art
// Takes a word (string) and the loaded banner lines
// Prints the word as ASCII art, 8 rows tall, with characters side by side
// This function handles the core logic of converting text to ASCII art
func printWord(word string, bannerLines []string) {
	// Loop through each of the 8 rows (height of each character)
	// i represents the current row number (0-7)
	for i := 0; i < charHeight; i++ {
		// For the current row, process each character in the word
		// This builds the ASCII art horizontally across the screen
		for _, char := range word {
			// Check if the character is within printable ASCII range
			// Printable ASCII codes: 32 (space) to 126 (tilde)
			// Characters outside this range (emojis, special Unicode) are ignored
			if char < 32 || char > 126 {
				continue // Skip this character and move to the next one
			}

			// Get the 8 lines of ASCII art for this character
			// This returns a slice of 8 strings, one for each row
			charArt := getCharArt(char, bannerLines)

			// Print the current row (i) of this character's ASCII art
			// charArt[i] gives the specific row for this character
			// Use fmt.Print (not Println) to keep characters on the same line
			fmt.Print(charArt[i])
		}
		// After finishing all characters for this row, print a newline
		// This moves to the next row of the ASCII art
		fmt.Println()
	}
}

// ========== MAIN FUNCTION - PROGRAM STARTS HERE ==========
// This is the entry point of the program - execution begins here
func main() {
	// Validate command line arguments
	// os.Args contains all command line arguments
	// os.Args[0] is the program name (e.g., "./program" or "go run .")
	// os.Args[1] is the first argument (the text to display)
	// os.Args[2] is the optional second argument (banner name)
	// We need at least 1 argument (the text), at most 2 arguments (text + banner)
	if len(os.Args) < 2 || len(os.Args) > 3 {
		// Print usage instructions to help the user
		// This is good practice - users need to know how to use your program
		fmt.Println("Usage: go run . \"text\" [standard|shadow|thinkertoy]")
		return // Exit the program - don't continue execution
	}

	// Get the input text from the first argument
	// Example: go run . "Hello" -> input = "Hello"
	input := os.Args[1]

	// Default banner file is "standard.txt"
	// This provides a sensible default if user doesn't specify a banner
	bannerFile := "standard.txt"

	// Optional banner selection - check if user provided a banner name
	if len(os.Args) == 3 {
		// Use a switch statement to handle different banner types
		// This is cleaner than multiple if-else statements
		switch os.Args[2] {
		case "standard":
			// Standard banner uses spaces, underscores, and pipes in a classic style
			bannerFile = "standard.txt"
		case "shadow":
			// Shadow banner creates a 3D shadow effect using different characters
			bannerFile = "shadow.txt"
		case "thinkertoy":
			// Thinkertoy banner uses an artistic style with different symbols
			bannerFile = "thinkertoy.txt"
		default:
			// If user provided an invalid banner name, show error and exit
			fmt.Println("Invalid banner type")
			return // Exit the program
		}
	}

	// Load the banner file into memory
	// loadBanner reads the file and returns the lines and any error
	bannerLines, err := loadBanner(bannerFile)

	// Check if there was an error loading the banner file
	if err != nil {
		// Print the error message (e.g., "file not found", "permission denied")
		fmt.Println("Error:", err)
		return // Exit the program - can't continue without banner
	}

	// Handle \n in input - split by literal backslash-n
	// The user might type "Hello\\nWorld" to mean two lines of text
	// strings.Split divides the input wherever it finds "\\n"
	// Example: "Hello\\nWorld" becomes ["Hello", "World"]
	words := strings.Split(input, "\\n")

	// Process each word/line from the split input
	// i is the index of the current word (0, 1, 2...)
	// word is the actual text content of the current line
	for i, word := range words {
		// Add a blank line between lines of text (except before the first line)
		// This creates visual separation when user wants multiple paragraphs
		// Example: "Hello\nWorld" should have a blank line between the two words
		if i > 0 {
			fmt.Println() // Print an extra blank line
		}

		// Skip empty lines - they don't need to be printed as ASCII art
		// Empty lines would result in no output anyway, so we can skip processing
		if word == "" {
			continue // Move to the next iteration of the loop
		}

		// Print the current word as ASCII art
		// This function handles the actual conversion from text to ASCII art
		printWord(word, bannerLines)
	}
}
