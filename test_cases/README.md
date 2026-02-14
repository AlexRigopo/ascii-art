# Test Cases Documentation

This directory contains comprehensive documentation of all test cases for the ASCII Art Generator project.

## Test Categories

1. **[Input Parsing Tests](input_tests.md)** - Tests for the ParseInput() function
2. **[Font Loading Tests](font_tests.md)** - Tests for font loading and glyph management
3. **[Rendering Tests](render_tests.md)** - Tests for ASCII art rendering
4. **[Edge Cases & Extreme Cases](edge_cases_extreme.md)** - Boundary conditions and stress tests

## Overview

- **Total Test Cases**: 31
- **Pass Rate**: 100%
- **Coverage**: Core functions (ParseInput, Render, LoadFont, HasGlyph, ParseFontLines)

## Test Execution

Run all tests:
```bash
go test ./tests -v
```

Run specific test file:
```bash
go test ./tests -v -run TestParseInput
go test ./tests -v -run TestFont
go test ./tests -v -run TestRender
```

## Test Status

✅ All tests passing
✅ No known bugs
✅ Full documentation available
