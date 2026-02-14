# ASCII Art Generator

A simple Go app that converts text into ASCII art using banner fonts. Just type some text and watch it turn into art.

## What It Does

- Takes text input and renders it as ASCII art
- Supports multiple lines with `\n` for line breaks
- Includes 3 different fonts: Standard, Shadow, and Thinkertoy
- Handles special characters (!, @, #, etc.)
- Normalizes line endings (Windows/Unix friendly)

## Quick Start

### Requirements

- Go 1.16+

### Build & Run

```bash
git clone <repository-url>
cd ascii-art
go build
./ascii-art "Hello World"
```

Or run directly:
```bash
go run . "Your text here"
```

## Usage Examples

### Basic Text
```bash
./ascii-art "Hello"
```
Outputs your text in ASCII art using the default (standard) font.

### Multiple Lines
```bash
./ascii-art "Hello\nWorld"
```
Use escaped newlines to render multiple lines.

### With Spaces & Special Characters
```bash
./ascii-art "Go-Lang 3.9!"
```
Fully supports spaces, numbers, punctuation, and special symbols.

## Project Structure

```
ascii-art/
├── main.go                        # Entry point
├── asciiart/                      # Core package
│   ├── font.go                   # Font loading & glyph management
│   ├── input.go                  # Parses input (handles \n)
│   ├── render.go                 # Renders ASCII art
│   └── errors.go                 # Error definitions
├── banners/                       # Font files
│   ├── standard.txt              # Default font
│   ├── shadow.txt                # Shadow style
│   └── thinkertoy.txt            # Thinkertoy style
├── tests/                         # Test suite
│   ├── input_test.go             # Input parsing tests
│   ├── font_test.go              # Font loading tests
│   └── render_test.go            # Rendering tests
└── test_cases/                    # Test documentation
    ├── input_tests.md
    ├── font_tests.md
    ├── render_tests.md
    └── TEST_EXECUTION_REPORT.md
```

## Testing

Run all 31 tests:
```bash
go test ./tests -v
```

Expected: All tests pass in ~5ms (Depending on the machine)

## Current State

✅ **Fully functional**
- Currently hardcoded to use the **standard** font
- No command-line flags for font selection (future enhancement)
- Works from the project root directory

### Why standard font only?

Kept it simple for now. The architecture supports multiple fonts — could add CLI flags if needed.

## Supported Characters

The application supports ASCII characters from space (32) to tilde (126), which includes:
- Letters (A-Z, a-z)
- Numbers (0-9)
- Punctuation and symbols

## API

### ParseInput

```go
func ParseInput(raw string) []string
```

Converts raw CLI input into logical lines, interpreting `\n` as newline characters.

### Render

```go
func Render(lines []string, font *Font) (string, error)
```

Renders text using the specified font and returns the ASCII art as a string.

### LoadFont

```go
func LoadFont(path string) (*Font, error)
```

Loads a banner font from a file.

## Examples

### Example 1: Simple Text

```bash
$ ./ascii-art "Hi"
 _   _     _
| | | |   (_)
| |_| | __ _
|  _  |/ _` |
| | | | (_| |
|_| |_|\__,_|
```

### Example 2: Multiple Lines

```bash
$ ./ascii-art "A\nB"
 _
/ \
\_/

 _
 < |
 > |
/_|
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

ASCII Art Generator - February 2026
