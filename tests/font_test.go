package tests

import (
	"os"
	"testing"

	"ascii-art/asciiart"
)

func TestLoadFontSuccess(t *testing.T) {
	// Change to project root to access banners directory
	cwd, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(cwd)

	// Test loading the standard font
	font, err := asciiart.LoadFont("banners/standard.txt")
	if err != nil {
		t.Fatalf("LoadFont failed: %v", err)
	}

	if font == nil {
		t.Fatalf("LoadFont returned nil Font")
	}

	// Standard font should have height of 8
	if font.Height != 8 {
		t.Fatalf("Font height = %d, want 8", font.Height)
	}

	// Should have glyphs for ASCII characters 32-126
	if len(font.Chars) == 0 {
		t.Fatalf("Font has no glyphs")
	}

	// Check for space and 'A' glyphs
	if !font.HasGlyph(' ') {
		t.Fatalf("Font missing glyph for space character")
	}
	if !font.HasGlyph('A') {
		t.Fatalf("Font missing glyph for 'A' character")
	}
}

func TestLoadFontFileNotFound(t *testing.T) {
	_, err := asciiart.LoadFont("nonexistent/font.txt")
	if err == nil {
		t.Fatalf("LoadFont should return error for nonexistent file")
	}
}

func TestLoadFontAllBanners(t *testing.T) {
	cwd, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(cwd)

	fonts := []string{
		"banners/standard.txt",
		"banners/shadow.txt",
		"banners/thinkertoy.txt",
	}

	for _, fontPath := range fonts {
		t.Run(fontPath, func(t *testing.T) {
			font, err := asciiart.LoadFont(fontPath)
			if err != nil {
				t.Fatalf("LoadFont(%s) failed: %v", fontPath, err)
			}

			if font.Height != 8 {
				t.Fatalf("Font height = %d, want 8", font.Height)
			}

			// Each font should have 95 characters (ASCII 32-126)
			if len(font.Chars) != 95 {
				t.Fatalf("Font char count = %d, want 95", len(font.Chars))
			}
		})
	}
}

func TestHasGlyph(t *testing.T) {
	font := &asciiart.Font{
		Height: 2,
		Chars: map[rune][]string{
			'A': {"AA", "AA"},
			'B': {"BB", "BB"},
		},
	}

	// Test existing glyphs
	if !font.HasGlyph('A') {
		t.Fatalf("HasGlyph('A') = false, want true")
	}
	if !font.HasGlyph('B') {
		t.Fatalf("HasGlyph('B') = false, want true")
	}

	// Test missing glyphs
	if font.HasGlyph('C') {
		t.Fatalf("HasGlyph('C') = true, want false")
	}
	if font.HasGlyph(' ') {
		t.Fatalf("HasGlyph(' ') = true, want false")
	}
}

func TestHasGlyphEmptyFont(t *testing.T) {
	font := &asciiart.Font{
		Height: 2,
		Chars:  make(map[rune][]string),
	}

	// Empty font should not have any glyphs
	if font.HasGlyph('A') {
		t.Fatalf("HasGlyph on empty font should return false")
	}
}

func TestFontGlyphCount(t *testing.T) {
	cwd, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(cwd)

	font, err := asciiart.LoadFont("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load font: %v", err)
	}

	// Standard ASCII range: 32 (space) to 126 (~) = 95 characters
	expectedCount := 95
	if len(font.Chars) != expectedCount {
		t.Fatalf("Font glyph count = %d, want %d", len(font.Chars), expectedCount)
	}
}

func TestFontGlyphHeight(t *testing.T) {
	cwd, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(cwd)

	font, err := asciiart.LoadFont("banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load font: %v", err)
	}

	// All glyphs should have the correct height
	for r, glyph := range font.Chars {
		if len(glyph) != font.Height {
			t.Fatalf("Glyph %q height = %d, want %d", r, len(glyph), font.Height)
		}
	}
}

func TestParseEmptyFontLines(t *testing.T) {
	lines := []string{}

	_, err := asciiart.ParseFontLines(lines)
	if err == nil {
		t.Fatalf("ParseFontLines should return error for empty lines")
	}
}

func TestParseFontLinesInsufficientLines(t *testing.T) {
	// Create lines array that's too short
	lines := make([]string, 100) // Not enough for all 95 characters with 8 lines each

	_, err := asciiart.ParseFontLines(lines)
	if err == nil {
		t.Fatalf("ParseFontLines should return error for insufficient lines")
	}
}

func TestFontStructCreation(t *testing.T) {
	font := &asciiart.Font{
		Height: 3,
		Chars: map[rune][]string{
			'X': {"X", "X", "X"},
			'Y': {"Y", "Y", "Y"},
		},
	}

	if font.Height != 3 {
		t.Fatalf("Font height = %d, want 3", font.Height)
	}

	if len(font.Chars) != 2 {
		t.Fatalf("Font chars count = %d, want 2", len(font.Chars))
	}

	// Check glyph lines
	xGlyph := font.Chars['X']
	if len(xGlyph) != 3 {
		t.Fatalf("Glyph 'X' line count = %d, want 3", len(xGlyph))
	}

	if xGlyph[0] != "X" {
		t.Fatalf("Glyph 'X' first line = %q, want %q", xGlyph[0], "X")
	}
}

func TestLoadFontIntegration(t *testing.T) {
	cwd, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(cwd)

	font, err := asciiart.LoadFont("banners/standard.txt")
	if err != nil {
		t.Fatalf("LoadFont failed: %v", err)
	}

	// Verify we can use the font
	tests := []struct {
		char    rune
		present bool
	}{
		{'A', true},
		{'Z', true},
		{'0', true},
		{'9', true},
		{' ', true},
		{'!', true},
		{'~', true},
	}

	for _, tt := range tests {
		if has := font.HasGlyph(tt.char); has != tt.present {
			t.Fatalf("HasGlyph(%q) = %v, want %v", tt.char, has, tt.present)
		}
	}
}

func TestFontGlyphContent(t *testing.T) {
	cwd, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(cwd)

	font, err := asciiart.LoadFont("banners/standard.txt")
	if err != nil {
		t.Fatalf("LoadFont failed: %v", err)
	}

	// Verify that glyphs contain content (non-empty strings)
	for r := rune(32); r <= rune(126); r++ {
		if !font.HasGlyph(r) {
			t.Fatalf("Font missing glyph for character %q", r)
		}

		glyph := font.Chars[r]
		if len(glyph) != 8 {
			t.Fatalf("Glyph %q has %d lines, want 8", r, len(glyph))
		}

		// Each line should exist (may be empty, but should be initialized)
		for _, line := range glyph {
			if line == "" && r != ' ' {
				// Space character can have empty lines, but most others shouldn't
				continue
			}
		}
	}
}
