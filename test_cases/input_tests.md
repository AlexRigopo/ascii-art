# Input Parsing Tests (ParseInput)

Function: `asciiart.ParseInput(raw string) []string`

Tests for the input parsing functionality that converts raw CLI strings into logical lines.

## Test Cases

### 1. TestParseInputNoNewline
**Purpose**: Test basic input without newlines  
**Input**: `"Hello"`  
**Expected Output**: `[]string{"Hello"}`  
**Category**: Basic Functionality  
**Status**: ✅ PASS

### 2. TestParseInputWithEscapedNewline
**Purpose**: Test escaped newline handling  
**Input**: `"Hello\nThere"`  
**Expected Output**: `[]string{"Hello", "There"}`  
**Category**: Core Functionality  
**Status**: ✅ PASS  
**Notes**: Interprets `\n` as actual newline character

### 3. TestParseInputLeadingAndTrailing
**Purpose**: Test leading and trailing newlines  
**Input**: `"\nHello\n"`  
**Expected Output**: `[]string{"", "Hello", ""}`  
**Category**: Edge Case  
**Status**: ✅ PASS  
**Notes**: Creates empty strings for blank lines

### 4. TestParseInputEmpty
**Purpose**: Test empty string input  
**Input**: `""`  
**Expected Output**: `[]string{""}`  
**Category**: Edge Case  
**Status**: ✅ PASS

### 5. TestParseInputSingleNewline
**Purpose**: Test single newline only  
**Input**: `"\n"`  
**Expected Output**: `[]string{"", ""}`  
**Category**: Edge Case  
**Status**: ✅ PASS

### 6. TestParseInputMultipleNewlines
**Purpose**: Test consecutive newlines  
**Input**: `"A\n\nB"`  
**Expected Output**: `[]string{"A", "", "B"}`  
**Category**: Edge Case  
**Status**: ✅ PASS

### 7. TestParseInputSpecialCharacters
**Purpose**: Test special characters in input  
**Input**: `"Hello!@#$"`  
**Expected Output**: `[]string{"Hello!@#$"}`  
**Category**: Special Characters  
**Status**: ✅ PASS

### 8. TestParseInputSpaces
**Purpose**: Test preservation of whitespace  
**Input**: `"  H  e  l  l  o  "`  
**Expected Output**: `[]string{"  H  e  l  l  o  "}`  
**Category**: Whitespace Handling  
**Status**: ✅ PASS  
**Notes**: Spaces are preserved exactly

### 9. TestParseInputMultipleLines
**Purpose**: Test multiple logical lines with newlines  
**Input**: `"Line1\nLine2\nLine3"`  
**Expected Output**: `[]string{"Line1", "Line2", "Line3"}`  
**Category**: Core Functionality  
**Status**: ✅ PASS

### 10. TestParseInputBackslashNotN
**Purpose**: Test non-newline escape sequences  
**Input**: `"Hello\tWorld"`  
**Expected Output**: `[]string{"Hello\tWorld"}`  
**Category**: Escape Sequence Handling  
**Status**: ✅ PASS  
**Notes**: Only `\n` is converted; `\t` stays literal

## Summary

| Category | Count | Pass |
|----------|-------|------|
| Basic | 1 | ✅ |
| Core | 2 | ✅ |
| Edge Case | 4 | ✅ |
| Special Chars | 1 | ✅ |
| Whitespace | 1 | ✅ |
| Escape Seq | 1 | ✅ |
| **Total** | **10** | **✅** |

## Key Behaviors Tested

✅ Escape sequence interpretation (`\n` as newline)  
✅ Empty string handling  
✅ Consecutive newlines  
✅ Whitespace preservation  
✅ Special character support  
✅ Non-newline escape sequences (literal)  
