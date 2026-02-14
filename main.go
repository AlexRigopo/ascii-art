package main

import (
	"fmt"
	"os"
	"path/filepath"

	"ascii-art/asciiart"
)

// main is the entry point for the ASCII Art Generator.
// It takes exactly one command-line argument (the text to render).
// Parses input, loads the standard font, renders & prints ASCII art.
// Exits with status 1 on error.
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . \"Your text here\"")
		os.Exit(1)
	}

	raw := os.Args[1]

	// Empty input → no output
	if raw == "" {
		return
	}

	lines := asciiart.ParseInput(raw)

	// Find font file relative to executable or working directory
	fontPath := findFontFile("banners/standard.txt")
	font, err := asciiart.LoadFont(fontPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading font:", err)
		os.Exit(1)
	}

	// Render and print
	output, err := asciiart.Render(lines, font)
	if err != nil {
		fmt.Fprintln(os.Stderr, "render error:", err)
		os.Exit(1)
	}

	fmt.Print(output)
}

// findFontFile looks for font in multiple paths
// First tries relative to executable, then relative to working directory
func findFontFile(relative string) string {
	// Try current working directory first (most common case)
	if _, err := os.Stat(relative); err == nil {
		return relative
	}

	// Try relative to executable path
	exe, err := os.Executable()
	if err == nil {
		exePath := filepath.Join(filepath.Dir(exe), relative)
		if _, err := os.Stat(exePath); err == nil {
			return exePath
		}
	}

	// Fall back to original (will fail with clear error)
	return relative
}
