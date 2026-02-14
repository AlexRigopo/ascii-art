# Font Loading & Management Tests

Functions: `LoadFont()`, `ParseFontLines()`, `HasGlyph()`

Tests for font file loading, parsing, and glyph management functionality.

## LoadFont() Tests

### 1. TestLoadFontSuccess
**Purpose**: Test successful font loading  
**Input**: Path to `"banners/standard.txt"`  
**Expected**: Non-nil Font struct with 95 glyphs at height 8  
**Status**: ✅ PASS  
**Key Assertions**:
- Font is not nil
- Font.Height == 8
- Font contains glyphs for space and 'A'
- len(Font.Chars) > 0

### 2. TestLoadFontFileNotFound
**Purpose**: Test error handling for missing files  
**Input**: Path to `"nonexistent/font.txt"`  
**Expected**: Error returned, nil Font  
**Status**: ✅ PASS  
**Key Assertions**:
- err != nil
- Error message contains file path info

### 3. TestLoadFontAllBanners
**Purpose**: Test loading all three banner fonts  
**Input**: 
- `"banners/standard.txt"`
- `"banners/shadow.txt"`
- `"banners/thinkertoy.txt"`  
**Expected**: All three load successfully with correct properties  
**Status**: ✅ PASS (subtests)  
**Subtests**:
- ✅ standard.txt: 95 glyphs, height 8
- ✅ shadow.txt: 95 glyphs, height 8
- ✅ thinkertoy.txt: 95 glyphs, height 8

## ParseFontLines() Tests

### 4. TestParseEmptyFontLines
**Purpose**: Test error handling for empty input  
**Input**: `[]string{}`  
**Expected**: Error returned  
**Status**: ✅ PASS

### 5. TestParseFontLinesInsufficientLines
**Purpose**: Test validation of line count  
**Input**: Array with only 100 lines (insufficient for 95 chars × 8 lines)  
**Expected**: Error returned  
**Status**: ✅ PASS  
**Notes**: Requires minimum 1 + (95 × 8) = 761 lines

## HasGlyph() Tests

### 6. TestHasGlyph
**Purpose**: Test glyph presence checking  
**Setup**: Font with glyphs for 'A' and 'B'  
**Tests**:
- `HasGlyph('A')` → true ✅
- `HasGlyph('B')` → true ✅
- `HasGlyph('C')` → false ✅
- `HasGlyph(' ')` → false ✅  
**Status**: ✅ PASS

### 7. TestHasGlyphEmptyFont
**Purpose**: Test empty font handling  
**Input**: Font with no glyphs (empty map)  
**Expected**: `HasGlyph()` returns false for any character  
**Status**: ✅ PASS

## Font Structure & Integration Tests

### 8. TestFontGlyphCount
**Purpose**: Verify correct glyph count  
**Input**: Loaded standard font  
**Expected**: 95 glyphs (ASCII 32-126)  
**Formula**: 126 - 32 + 1 = 95  
**Status**: ✅ PASS

### 9. TestFontGlyphHeight
**Purpose**: Verify all glyphs have correct height  
**Input**: Loaded standard font  
**Expected**: All 95 glyphs have 8 lines  
**Status**: ✅ PASS

### 10. TestFontStructCreation
**Purpose**: Test Font struct instantiation  
**Setup**:
```go
Font{
  Height: 3,
  Chars: map[rune][]string{
    'X': {"X", "X", "X"},
    'Y': {"Y", "Y", "Y"},
  }
}
```
**Expected**: Struct created correctly  
**Status**: ✅ PASS

### 11. TestLoadFontIntegration
**Purpose**: Full integration test with real font  
**Input**: Loaded standard font  
**Tests Coverage**:
- ASCII letters (A-Z)
- Digits (0-9)
- Whitespace
- Punctuation (!, ~)  
**Status**: ✅ PASS

### 12. TestFontGlyphContent
**Purpose**: Verify glyph content integrity  
**Input**: Loaded standard font  
**Expected**: 
- All 95 characters have glyphs
- Each glyph has exactly 8 lines
- Lines are initialized (content varies)  
**Status**: ✅ PASS

## Summary

| Category | Count | Pass |
|----------|-------|------|
| LoadFont | 3 | ✅ |
| ParseFontLines | 2 | ✅ |
| HasGlyph | 2 | ✅ |
| Structure & Integration | 5 | ✅ |
| **Total** | **12** | **✅** |

## Key Behaviors Tested

✅ Font file loading and parsing  
✅ Error handling (missing files, insufficient data)  
✅ Glyph presence checking  
✅ All three banner fonts (standard, shadow, thinkertoy)  
✅ ASCII character range coverage (32-126)  
✅ Glyph height validation  
✅ Font structure initialization
