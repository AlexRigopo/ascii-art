package asciiart

import "strings"

// ParseInput converts the CLI string into logical lines.
// Interprets \n as newlines and splits accordingly.
// Other escapes (like \t) are kept literal.
func ParseInput(raw string) []string {
	var b strings.Builder

	for i := 0; i < len(raw); i++ {
		if raw[i] == '\\' && i+1 < len(raw) {
			switch raw[i+1] {
			case 'n':
				b.WriteByte('\n')
				i++ // skip 'n'
				continue
			default:
				// Keep the backslash literally if it's not "\n"
				b.WriteByte(raw[i])
			}
		} else {
			b.WriteByte(raw[i])
		}
	}

	normalized := b.String()
	// Normalize Windows-style newlines just in case
	normalized = strings.ReplaceAll(normalized, "\r\n", "\n")

	// Split into logical lines. Example:
	// "Hello\nThere" -> []{"Hello", "There"}
	return strings.Split(normalized, "\n")
}
