package main

import (
	"fmt"
	"os"

	"ascii-art/asciiart"
)

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

	// For now, always use standard.txt; later you can add flags for shadow/thinkertoy.
	font, err := asciiart.LoadFont("standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error loading font:", err)
		os.Exit(1)
	}

	output, err := asciiart.Render(lines, font)
	if err != nil {
		fmt.Fprintln(os.Stderr, "render error:", err)
		os.Exit(1)
	}

	fmt.Print(output)
}
