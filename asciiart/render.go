package asciiart

import (
	"fmt"
	"strings"
)

// Render converts text lines into ASCII art using the font.
// Renders each non-empty line at font.Height rows.
// Empty lines become single newlines.
// Returns error if a character isn't supported or glyph is missing.
func Render(lines []string, font *Font) (string, error) {
	var out strings.Builder

	for i, line := range lines {
		// Empty line -> just newline
		if line == "" {
			out.WriteByte('\n')
			continue
		}

		// Each row of the font
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
