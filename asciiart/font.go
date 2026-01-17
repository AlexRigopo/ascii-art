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

	font, err := parseFontLines(lines)
	if err != nil {
		return nil, fmt.Errorf("parse font: %w", err)
	}
	return font, nil
}

// parseFontLines builds a Font from the file lines.
// Separated to allow easier testing if you want to later.
func parseFontLines(lines []string) (*Font, error) {
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

// HasGlyph reports whether the font contains this rune.
func (f *Font) HasGlyph(r rune) bool {
	_, ok := f.Chars[r]
	return ok
}
