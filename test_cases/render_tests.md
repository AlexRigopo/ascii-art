# ASCII Art Rendering Tests (Render)

Function: `asciiart.Render(lines []string, font *Font) (string, error)`

Tests for converting text lines into ASCII art using font glyphs.

## Basic Rendering Tests

### 1. TestRenderSingleLine
**Purpose**: Test basic single-line rendering  
**Input**:
- lines: `[]string{"AB"}`
- font: Test font (height 2) with glyphs: A={'AA','AA'}, B={'BB','BB'}  
**Expected Output**:
```
AABB
AABB

```
**Status**: ✅ PASS

### 2. TestRenderWithSpace
**Purpose**: Test rendering with space character  
**Input**:
- lines: `[]string{"A B"}`
- font: Test font with space={'__','__'}  
**Expected Output**:
```
AA__BB
AA__BB

```
**Status**: ✅ PASS

### 3. TestRenderSingleCharacter
**Purpose**: Test rendering single character  
**Input**:
- lines: `[]string{"A"}`  
**Expected Output**:
```
AA
AA

```
**Status**: ✅ PASS

### 4. TestRenderOnlySpaces
**Purpose**: Test rendering string of spaces  
**Input**:
- lines: `[]string{"   "}` (3 spaces)  
**Expected Output**:
```
______
______

```
**Status**: ✅ PASS  
**Notes**: Each space renders as '__'

## Multi-Line Tests

### 5. TestRenderMultipleLogicalLines
**Purpose**: Test rendering multiple logical lines with blank lines  
**Input**:
```
lines: []string{"AB", "", "BA"}
```
**Expected Output**:
```
AABB
AABB

BBAA
BBAA

```
**Status**: ✅ PASS  
**Notes**: Empty string produces single newline

### 6. TestRenderBlankLine
**Purpose**: Test rendering single blank line  
**Input**: `[]string{""}`  
**Expected Output**: `"\n"` (just one newline)  
**Status**: ✅ PASS

## Edge Case Tests

### 7. TestRenderEmptyLineList
**Purpose**: Test empty input  
**Input**: `[]string{}`  
**Expected Output**: `""` (empty string)  
**Status**: ✅ PASS

## Error Handling Tests

### 8. TestRenderUnsupportedCharacter
**Purpose**: Test error for unsupported character  
**Input**: 
- lines: `[]string{"€"}` (Euro symbol, code 8364)
- font: Standard ASCII font (32-126)  
**Expected**: 
- Error returned
- Error message: `"unsupported character '€' (code 8364)"`  
**Status**: ✅ PASS  
**Error Type**: Out-of-range character

### 9. TestRenderMissingGlyph
**Purpose**: Test error for missing glyph  
**Input**:
- lines: `[]string{"AB"}`
- font: Font with 'A' but missing 'B'  
**Expected**:
- Error returned
- Error message: `"missing glyph for 'B'"`  
**Status**: ✅ PASS  
**Error Type**: Incomplete font data

## Summary

| Category | Count | Pass |
|----------|-------|------|
| Basic Rendering | 4 | ✅ |
| Multi-Line | 2 | ✅ |
| Edge Cases | 1 | ✅ |
| Error Handling | 2 | ✅ |
| **Total** | **9** | **✅** |

## Key Behaviors Tested

✅ Single and multiple lines  
✅ Space character handling  
✅ Blank line handling (produces newline)  
✅ Correct ASCII art alignment (stacked glyphs)  
✅ Error handling for unsupported characters  
✅ Error handling for missing glyphs  
✅ Output format validation (final newline on each row)

## Output Format Rules

1. Each logical line is rendered at font.Height vertical rows
2. Each row ends with newline character (`\n`)
3. Empty logical lines produce single newline
4. No extra separator lines between logical lines
5. Final output includes trailing newline from last row
