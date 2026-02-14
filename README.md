# ASCII Art Generator

A Go application that converts text input into ASCII art using various banner fonts.

## Features

- Multiple font styles (Standard, Shadow, Thinkertoy)
- Support for escaped newlines (`\n`)
- Customizable font loading
- Comprehensive test coverage

## Installation

### Prerequisites

- Go 1.16 or higher

### Build

```bash
git clone <repository-url>
cd ascii-art
go build -o ascii-art main.go
```

## Usage

### Basic Usage

```bash
./ascii-art "Hello World"
```

### With Different Font

```bash
./ascii-art --font=shadow "Hello"
./ascii-art --font=thinkertoy "World"
```

### Multiple Lines

Use `\n` for line breaks:

```bash
./ascii-art "Hello\nWorld"
```

## Project Structure

```
.
├── main.go                 # Entry point
├── asciiart/
│   ├── font.go            # Font loading and management
│   ├── input.go           # Input parsing
│   ├── render.go          # ASCII art rendering
│   └── errors.go          # Error handling
├── banners/
│   ├── standard.txt       # Standard font
│   ├── shadow.txt         # Shadow font
│   └── thinkertoy.txt     # Thinkertoy font
├── tests/
│   ├── input_test.go      # Input parsing tests
│   └── render_test.go     # Rendering tests
├── go.mod                 # Go module definition
└── README.md              # This file
```

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test ./tests -v
```

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
