package asciiart

import (
	"bufio"
	"fmt"
	"os"
)

const (
	firstRune  = 32 // ' '
	lastRune   = 126
	fontHeight = 8
)

type Font struct {
	Height int
	Chars  map[rune][]string // rune -> [Height]lines
}

// LoadFont loads a banner font from a file path, using the standard
// 8-line per glyph format used by the project subjects.
//
// Parameters:
//   - path: The file path to the banner font file (e.g., "banners/standard.txt").
//
// The font file must contain glyph definitions for all ASCII characters from
// space (32) to tilde (126), each glyph being 8 lines tall.
//
// Returns:
//
//	*Font: A pointer to the loaded Font struct containing all glyphs.
//	error: An error if the file cannot be opened, read, or parsed.
func LoadFont(path string) (*Font, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open font file: %w", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Keep lines exactly as-is (no trimming)
		line := scanner.Text()
		// Normalize possible CRLF -> LF already removed by Scanner
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan font file: %w", err)
	}

	font, err := ParseFontLines(lines)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}
	return font, nil
}

// ParseFontLines builds a Font struct from parsed file lines.
//
// Parameters:
//   - lines: A slice of strings representing lines read from the font file.
//
// This helper function processes the raw file lines into a Font struct with
// glyph definitions for all ASCII characters (space to tilde).
// Each character glyph is expected to be exactly fontHeight (8) lines.
// Separator lines between glyphs are automatically skipped.
//
// Returns:
//
//	*Font: A pointer to the fully constructed Font struct.
//	error: An error if the file is malformed or contains insufficient lines.
func ParseFontLines(lines []string) (*Font, error) {
	charCount := lastRune - firstRune + 1

	// Minimum lines:
	// 1 header + charCount * fontHeight (glyph lines)
	minLines := 1 + charCount*fontHeight
	if len(lines) < minLines {
		return nil, fmt.Errorf("font file too short: got %d lines, want at least %d", len(lines), minLines)
	}

	font := &Font{
		Height: fontHeight,
		Chars:  make(map[rune][]string, charCount),
	}

	idx := 1 // skip header line

	for r := firstRune; r <= lastRune; r++ {
		// Read the 8 glyph lines
		glyph := make([]string, fontHeight)
		for row := 0; row < fontHeight; row++ {
			if idx >= len(lines) {
				return nil, fmt.Errorf("unexpected end of font at rune %d row %d", r, row)
			}
			glyph[row] = lines[idx]
			idx++
		}

		// Skip a separator line *if* we're not at the last rune
		// and there's still a line left.
		if r != lastRune && idx < len(lines) {
			// separator line (often "")
			idx++
		}

		font.Chars[rune(r)] = glyph
	}

	return font, nil
}

// HasGlyph checks if the font contains a glyph definition for the given character.
//
// Parameters:
//   - r: The rune (character) to check for in the font.
//
// Returns:
//
//	bool: true if the font has a glyph for this character, false otherwise.
func (f *Font) HasGlyph(r rune) bool {
	_, ok := f.Chars[r]
	return ok
}
