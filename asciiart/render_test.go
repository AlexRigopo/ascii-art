package asciiart

import "testing"

func makeTestFont() *Font {
	// height 2 so we can write expected strings easily
	return &Font{
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

	got, err := Render(lines, font)
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

	got, err := Render(lines, font)
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

	got, err := Render(lines, font)
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
