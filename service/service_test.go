package service

import (
	"testing"
	"unicode"
)

func TestGenerate6Alphanumeric(t *testing.T) {
    // Test that generated string is exactly 6 characters long
    s := generate6Alphanumeric()
    if len(s) != 6 {
        t.Errorf("generate6Alphanumeric() returned string with length %d, expected 6", len(s))
    }
    
    // Test that generated string contains only alphanumeric characters
    for _, c := range s {
        if !unicode.IsDigit(c) && !unicode.IsLetter(c) {
            t.Errorf("generate6Alphanumeric() returned string %s with non-alphanumeric character %c", s, c)
        }
    }
    
    // Test that generated string contains at least one letter and one digit
    hasLetter := false
    hasDigit := false
    for _, c := range s {
        if unicode.IsLetter(c) {
            hasLetter = true
        } else if unicode.IsDigit(c) {
            hasDigit = true
        }
    }
    if !hasLetter || !hasDigit {
        t.Errorf("generate6Alphanumeric() returned string %s without at least one letter and one digit", s)
    }
}


