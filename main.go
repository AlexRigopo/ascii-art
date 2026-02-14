package main

import (
	"fmt"
	"os"

	"ascii-art/asciiart"
)

// main is the entry point for the ASCII Art Generator.
// It takes exactly one command-line argument (the text to render as ASCII art).
// It parses the input, loads the standard banner font, renders the text,
// and prints the result to standard output.
// Exits with status code 1 if arguments are invalid or rendering fails.
func main() {
	if len(os.Args) != 2 {
		// The subject examples assume exactly one argument.
		// You can tweak this if your evaluator wants silent failure instead.
		fmt.Fprintln(os.Stderr, "Usage: go run . \"Your text here\"")
		os.Exit(1)
	}

	raw := os.Args[1]

	// Special case: empty string → no output at all (see subject example).
	if raw == "" {
		return
	}

	lines := asciiart.ParseInput(raw)

	// Load the standard banner font (default)
	font, err := asciiart.LoadFont("banners/standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading font:", err)
		os.Exit(1)
	}

	// Render the ASCII art
	output, err := asciiart.Render(lines, font)
	if err != nil {
		fmt.Fprintln(os.Stderr, "render error:", err)
		os.Exit(1)
	}

	fmt.Print(output)
}
