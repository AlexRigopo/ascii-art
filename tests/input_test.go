package tests

import (
	"reflect"
	"testing"

	"ascii-art/asciiart"
)

func TestParseInputNoNewline(t *testing.T) {
	got := asciiart.ParseInput("Hello")
	want := []string{"Hello"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"Hello\") = %#v, want %#v", got, want)
	}
}

func TestParseInputWithEscapedNewline(t *testing.T) {
	got := asciiart.ParseInput("Hello\\nThere")
	want := []string{"Hello", "There"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"Hello\\\\nThere\") = %#v, want %#v", got, want)
	}
}

func TestParseInputLeadingAndTrailing(t *testing.T) {
	got := asciiart.ParseInput("\\nHello\\n")
	want := []string{"", "Hello", ""}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"\\\\nHello\\\\n\") = %#v, want %#v", got, want)
	}
}

func TestParseInputEmpty(t *testing.T) {
	got := asciiart.ParseInput("")
	want := []string{""}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"\") = %#v, want %#v", got, want)
	}
}

func TestParseInputSingleNewline(t *testing.T) {
	got := asciiart.ParseInput("\\n")
	want := []string{"", ""}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"\\\\n\") = %#v, want %#v", got, want)
	}
}

func TestParseInputMultipleNewlines(t *testing.T) {
	got := asciiart.ParseInput("A\\n\\nB")
	want := []string{"A", "", "B"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"A\\\\n\\\\nB\") = %#v, want %#v", got, want)
	}
}

func TestParseInputSpecialCharacters(t *testing.T) {
	got := asciiart.ParseInput("Hello!@#$")
	want := []string{"Hello!@#$"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"Hello!@#$\") = %#v, want %#v", got, want)
	}
}

func TestParseInputSpaces(t *testing.T) {
	got := asciiart.ParseInput("  H  e  l  l  o  ")
	want := []string{"  H  e  l  l  o  "}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput with spaces = %#v, want %#v", got, want)
	}
}

func TestParseInputMultipleLines(t *testing.T) {
	got := asciiart.ParseInput("Line1\\nLine2\\nLine3")
	want := []string{"Line1", "Line2", "Line3"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput with multiple lines = %#v, want %#v", got, want)
	}
}

func TestParseInputBackslashNotN(t *testing.T) {
	got := asciiart.ParseInput("Hello\\tWorld")
	want := []string{"Hello\\tWorld"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"Hello\\\\tWorld\") = %#v, want %#v", got, want)
	}
}
