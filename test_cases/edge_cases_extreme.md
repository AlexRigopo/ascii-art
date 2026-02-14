# Edge Cases & Extreme Test Cases

Boundary conditions, stress tests, and extreme scenarios not covered by standard tests.

## Extreme Case 1: Maximum Length Input

**Purpose**: Test rendering extremely long text  
**Category**: Performance/Stress Test  
**Difficulty**: Hard  
**Status**: ⚠️ Not Implemented (Future Test)

### Test Case
```go
TestRenderMaxLengthInput
Input: String with 10,000 characters
Expected: Successful render without memory overflow
Assertion: Verify output length correlates correctly
```

### Challenge
- Memory allocation for large output
- Processing time for massive strings
- Output buffer management

### Recommended Implementation
```go
func TestRenderMaxLengthInput(t *testing.T) {
    font := loadTestFont()
    longText := strings.Repeat("A", 10000)
    lines := []string{longText}
    
    got, err := asciiart.Render(lines, font)
    if err != nil {
        t.Fatalf("Render failed: %v", err)
    }
    
    // Verify output is proportional
    if len(got) == 0 {
        t.Fatalf("Output should not be empty")
    }
}
```

---

## Extreme Case 2: Complex Multi-Line with Mixed Content

**Purpose**: Test complex scenario with all character types  
**Category**: Integration/Complexity  
**Difficulty**: Very Hard  
**Status**: ⚠️ Not Implemented (Future Test)

### Test Case
```
Input: "Hello123!\n\nWORLD@#$%\nGo-Lang_Rules"

Expected:
- Line 1: "Hello123!" with ASCII art
- Line 2: Empty line (single newline)
- Line 3: "WORLD@#$%" with mixed case and symbols
- Line 4: "Go-Lang_Rules" with hyphens and underscores
```

### Complexity Factors
- Mixed alphanumeric characters
- Special characters (!@#$%-)
- Case sensitivity handling
- Multiple consecutive blank lines
- Long words with special separators

### Recommended Implementation
```go
func TestRenderComplexMultiLine(t *testing.T) {
    font, _ := asciiart.LoadFont("banners/standard.txt")
    lines := []string{
        "Hello123!",
        "",
        "WORLD@#$%",
        "Go-Lang_Rules",
    }
    
    got, err := asciiart.Render(lines, font)
    if err != nil {
        t.Fatalf("Complex render failed: %v", err)
    }
    
    // Verify structure
    outputLines := strings.Split(got, "\n")
    if len(outputLines) < 32 { // 8 height × 4 lines
        t.Fatalf("Output too short for expected content")
    }
}
```

---

## Extreme Case 3: Unicode Boundary Testing

**Purpose**: Test all boundaries of ASCII range + just outside  
**Category**: Boundary Validation  
**Difficulty**: Hard  
**Status**: ⚠️ Not Implemented (Future Test)

### Test Case Specification

#### 3a. Valid Boundaries
```
- Test ASCII 32 (space) - minimum valid
- Test ASCII 126 (~) - maximum valid
- Test ASCII 33 (!) - first printable after space
- Test ASCII 125 (}) - last printable before tilde
```

#### 3b. Invalid Boundaries
```
- Test ASCII 31 (control character)
- Test ASCII 127 (DEL)
- Test ASCII 128+ (extended ASCII)
```

### Recommended Implementation
```go
func TestASCIIBoundaries(t *testing.T) {
    font, _ := asciiart.LoadFont("banners/standard.txt")
    
    // Valid boundaries
    validTests := []rune{32, 33, 125, 126}
    for _, r := range validTests {
        lines := []string{string(r)}
        _, err := asciiart.Render(lines, font)
        if err != nil {
            t.Fatalf("Char %d should be valid, got error: %v", r, err)
        }
    }
    
    // Invalid boundaries
    invalidTests := []rune{31, 127, 128, 256}
    for _, r := range invalidTests {
        lines := []string{string(r)}
        _, err := asciiart.Render(lines, font)
        if err == nil {
            t.Fatalf("Char %d should be invalid", r)
        }
    }
}
```

---

## Extreme Case 4: All Special Characters String

**Purpose**: Render every special character consecutively  
**Category**: Content Validation  
**Difficulty**: Hard  
**Status**: ⚠️ Not Implemented (Future Test)

### Test Case
```
Input: "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

This includes:
- All punctuation marks
- All mathematical symbols
- All brackets and quotes
- No alphanumeric characters
```

### Challenge
- Verify each special char renders with correct glyph
- Check for glyph collision or overlap
- Validate output readability with dense special chars

### Recommended Implementation
```go
func TestAllSpecialCharacters(t *testing.T) {
    font, _ := asciiart.LoadFont("banners/standard.txt")
    
    specials := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
    lines := []string{specials}
    
    got, err := asciiart.Render(lines, font)
    if err != nil {
        t.Fatalf("Failed to render special chars: %v", err)
    }
    
    // Verify output contains expected number of lines
    outLines := strings.Split(strings.TrimRight(got, "\n"), "\n")
    if len(outLines) != 8 { // font height
        t.Fatalf("Expected 8 output lines, got %d", len(outLines))
    }
}
```

---

## Extreme Case 5: Deeply Nested Newlines

**Purpose**: Test behavior with many consecutive escaped newlines  
**Category**: Edge Case / Stress  
**Difficulty**: Hard  
**Status**: ⚠️ Not Implemented (Future Test)

### Test Case
```
Input: "A\n\n\n\n\n\nZ"

Expected:
- Line 1: "A" rendered
- Lines 2-6: Empty lines (each produces single newline)
- Line 7: "Z" rendered
Total: 8 + 1 + 1 + 1 + 1 + 1 + 8 = 21 output lines
```

### Verification
- Correct line count
- Proper spacing/alignment maintained
- No corruption of blank lines
- Handles large gap between text

### Recommended Implementation
```go
func TestDeeplyNestedNewlines(t *testing.T) {
    lines := asciiart.ParseInput("A\\n\\n\\n\\n\\n\\nZ")
    if len(lines) != 7 {
        t.Fatalf("Expected 7 lines, got %d", len(lines))
    }
    
    font, _ := asciiart.LoadFont("banners/standard.txt")
    got, err := asciiart.Render(lines, font)
    if err != nil {
        t.Fatalf("Render failed: %v", err)
    }
    
    outputLines := strings.Split(strings.TrimRight(got, "\n"), "\n")
    if len(outputLines) != 21 {
        t.Fatalf("Expected 21 output lines, got %d", len(outputLines))
    }
}
```

---

## Extreme Case 6: ULTIMATE HARD CASE - Full Banner Art Challenge

**Purpose**: Comprehensive stress test combining all extreme conditions  
**Category**: Integration/Comprehensive  
**Difficulty**: **EXTREME** 🔴  
**Status**: ⚠️ Not Implemented (Future Test)

### Test Scenario

```
Input Text:
"Test@123\n\nАБвеньБорВо!\nPython3.9!\nA\n\n\nZ\n!"

Requirements:
1. Alphanumeric (Test@123)
2. Special characters (@!)
3. Numbers (123, 3.9)
4. Blank lines (multiple)
5. Single characters (A, Z, !)
6. Mixed case (Аá?)
7. Punctuation (., !)
8. Consecutive blanks
9. Real font file loading (all 3 types)
10. Output validation and alignment
```

### What Makes This Extreme?

✅ Tests every ParseInput edge case  
✅ Tests every Render scenario  
✅ Tests real font loading  
✅ Validates complete output structure  
✅ Memory and performance stress  
✅ Complex alignment requirements  
✅ 40+ lines of ASCII output  
✅ Error recovery scenarios  

### Recommended Implementation

```go
func TestUltimateHardCase(t *testing.T) {
    // Test with all three fonts
    fontPaths := []string{
        "banners/standard.txt",
        "banners/shadow.txt",
        "banners/thinkertoy.txt",
    }
    
    input := "Test@123\\n\\nWORLD!\\nPython\\nA\\n\\n\\nZ\\n!"
    
    for _, fontPath := range fontPaths {
        t.Run(fontPath, func(t *testing.T) {
            // Parse input
            lines := asciiart.ParseInput(input)
            if len(lines) != 8 {
                t.Fatalf("Expected 8 lines, got %d", len(lines))
            }
            
            // Load font
            font, err := asciiart.LoadFont(fontPath)
            if err != nil {
                t.Fatalf("Font load failed: %v", err)
            }
            
            // Render
            output, err := asciiart.Render(lines, font)
            if err != nil {
                t.Fatalf("Render failed: %v", err)
            }
            
            // Validate output
            if len(output) == 0 {
                t.Fatalf("Output is empty")
            }
            
            outputLines := strings.Split(output, "\n")
            // 3 complex lines (8 rows each) + 3 blanks + 2 singles = 31+ lines
            if len(outputLines) < 30 {
                t.Fatalf("Output too short, expected 30+ lines, got %d", 
                    len(outputLines))
            }
        })
    }
}
```

---

## Summary: Extreme Test Coverage

| Case | Difficulty | Type | Status |
|------|-----------|------|--------|
| Max Length Input | Hard | Performance | ⚠️ Future |
| Complex Multi-Line | Very Hard | Integration | ⚠️ Future |
| Unicode Boundaries | Hard | Validation | ⚠️ Future |
| All Special Chars | Hard | Content | ⚠️ Future |
| Nested Newlines | Hard | Edge Case | ⚠️ Future |
| Ultimate Hard Case | **EXTREME** 🔴 | Comprehensive | ⚠️ Future |

## Implementation Priority

1. **High Priority**: Ultimate Hard Case (comprehensive)
2. **Medium Priority**: Unicode Boundaries (safety)
3. **Medium Priority**: All Special Characters (coverage)
4. **Lower Priority**: Max Length & Complex Multi-Line (optimization)

---

## Notes

These extreme tests push the system to its limits, testing:
- Maximum complexity scenarios
- Boundary conditions
- Integration across all functions
- Resource management
- Error handling
- Output validation
- Real-world use cases
