package tests

import (
	"testing"

	"ascii-art/asciiart"
)

func makeTestFont() *asciiart.Font {
	// height 2 so we can write expected strings easily
	return &asciiart.Font{
		Height: 2,
		Chars: map[rune][]string{
			' ': {"__", "__"},
			'A': {"AA", "AA"},
			'B': {"BB", "BB"},
		},
	}
}

func TestRenderSingleLine(t *testing.T) {
	font := makeTestFont()
	lines := []string{"AB"}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	// Expected:
	// Row 0: "AABB"
	// Row 1: "AABB"
	// Plus final newline at each row.
	want := "AABB\nAABB\n"

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}

func TestRenderWithSpace(t *testing.T) {
	font := makeTestFont()
	lines := []string{"A B"}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	// Row 0: "AA__BB"
	// Row 1: "AA__BB"
	want := "AA__BB\nAA__BB\n"

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}

func TestRenderMultipleLogicalLines(t *testing.T) {
	font := makeTestFont()
	lines := []string{"AB", "", "BA"}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	// "AB" -> 2 rows
	// ""   -> one blank line
	// "BA" -> 2 rows
	want := "" +
		"AABB\n" +
		"AABB\n" +
		"\n" +
		"BBAA\n" +
		"BBAA\n"

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}

func TestRenderEmptyLineList(t *testing.T) {
	font := makeTestFont()
	lines := []string{}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	// Empty input should produce empty output
	want := ""

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}

func TestRenderBlankLine(t *testing.T) {
	font := makeTestFont()
	lines := []string{""}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	// Single blank line should produce just a newline
	want := "\n"

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}

func TestRenderUnsupportedCharacter(t *testing.T) {
	font := makeTestFont()
	lines := []string{"€"} // EUR symbol, not in ASCII 32-126

	_, err := asciiart.Render(lines, font)
	if err == nil {
		t.Fatalf("Render should return error for unsupported character")
	}

	if err.Error() != "unsupported character '€' (code 8364)" {
		t.Fatalf("Render error message incorrect: %v", err)
	}
}

func TestRenderMissingGlyph(t *testing.T) {
	font := &asciiart.Font{
		Height: 2,
		Chars: map[rune][]string{
			'A': {"AA", "AA"},
			// 'B' is missing
		},
	}
	lines := []string{"AB"}

	_, err := asciiart.Render(lines, font)
	if err == nil {
		t.Fatalf("Render should return error for missing glyph")
	}

	if err.Error() != "missing glyph for 'B'" {
		t.Fatalf("Render error message incorrect: want 'missing glyph for 'B'', got: %v", err)
	}
}

func TestRenderSingleCharacter(t *testing.T) {
	font := makeTestFont()
	lines := []string{"A"}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	want := "AA\nAA\n"

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}

func TestRenderOnlySpaces(t *testing.T) {
	font := makeTestFont()
	lines := []string{"   "}

	got, err := asciiart.Render(lines, font)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	// Three spaces should render as "______" on each row
	want := "______\n______\n"

	if got != want {
		t.Fatalf("Render(%v) = %q, want %q", lines, got, want)
	}
}
