package asciiart

import (
	"fmt"
	"strings"
)

// Render converts text lines into ASCII art using the provided font.
//
// Parameters:
//   - lines: A slice of text strings to render as ASCII art.
//   - font: A pointer to a Font struct containing glyph definitions.
//
// For each non-empty line, it renders the text at font.Height vertical rows.
// Empty lines produce a single newline. Unsupported characters return an error.
// The function does not print; it returns the complete ASCII art as a string.
//
// Returns:
//
//	string: The complete ASCII art with newlines separating rows.
//	error: An error if an unsupported character is encountered or a glyph is missing.
func Render(lines []string, font *Font) (string, error) {
	var out strings.Builder

	for i, line := range lines {
		// Empty logical line -> just a newline in output
		if line == "" {
			out.WriteByte('\n')
			continue
		}

		// For each row of the font
		for row := 0; row < font.Height; row++ {
			for _, r := range line {
				if r < firstRune || r > lastRune {
					return "", fmt.Errorf("unsupported character %q (code %d)", r, r)
				}
				glyph, ok := font.Chars[r]
				if !ok {
					return "", fmt.Errorf("missing glyph for %q", r)
				}
				out.WriteString(glyph[row])
			}
			out.WriteByte('\n')
		}

		// No extra blank line between logical lines:
		// the next iteration will either print ASCII art or a single '\n'.
		_ = i // kept in case you want more control per line later
	}

	return out.String(), nil
}
