package asciiart

import "errors"

var (
	ErrUnsupportedRune = errors.New("unsupported rune")
	ErrMissingGlyph    = errors.New("missing glyph in font")
)
