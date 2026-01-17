package asciiart

import (
	"reflect"
	"testing"
)

func TestParseInputNoNewline(t *testing.T) {
	got := ParseInput("Hello")
	want := []string{"Hello"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"Hello\") = %#v, want %#v", got, want)
	}
}

func TestParseInputWithEscapedNewline(t *testing.T) {
	got := ParseInput("Hello\\nThere")
	want := []string{"Hello", "There"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"Hello\\\\nThere\") = %#v, want %#v", got, want)
	}
}

func TestParseInputLeadingAndTrailing(t *testing.T) {
	got := ParseInput("\\nHello\\n")
	want := []string{"", "Hello", ""}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseInput(\"\\\\nHello\\\\n\") = %#v, want %#v", got, want)
	}
}
