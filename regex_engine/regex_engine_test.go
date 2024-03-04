package regexengine_test

import (
	regexengine "softwaredesign/regex_engine"
	"testing"
)

func TestMatch(t *testing.T) {
	// Test case: Matching with '^' pattern
	pattern := "^ab"
	text := "abcdef"
	expected := true
	result := regexengine.Match(pattern, text)
	if result != expected {
		t.Errorf("match(%q, %q) = %t, expected %t", pattern, text, result, expected)
	}

	// Test case: Matching without '^' pattern
	pattern = "def$"
	text = "abcdef"
	expected = true
	result = regexengine.Match(pattern, text)
	if result != expected {
		t.Errorf("match(%q, %q) = %t, expected %t", pattern, text, result, expected)
	}

	// Test case: Matching with '*' pattern
	pattern = "a*bc"
	text = "abbbbbbc"
	expected = true
	result = regexengine.Match(pattern, text)
	if result != expected {
		t.Errorf("match(%q, %q) = %t, expected %t", pattern, text, result, expected)
	}

	// Test case: Matching with '.' pattern
	pattern = "a.c"
	text = "abc"
	expected = true
	result = regexengine.Match(pattern, text)
	if result != expected {
		t.Errorf("match(%q, %q) = %t, expected %t", pattern, text, result, expected)
	}

	// Test case: No match
	pattern = "xyz"
	text = "abc"
	expected = false
	result = regexengine.Match(pattern, text)
	if result != expected {
		t.Errorf("match(%q, %q) = %t, expected %t", pattern, text, result, expected)
	}
}
