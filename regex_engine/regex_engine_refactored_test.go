package regexengine_test

import (
	regexengine "softwaredesign/regex_engine"
	"testing"
)

func TestRxBase_Match(t *testing.T) {

	// MockLitMatcher with matching chars
	matcher := &regexengine.RxLit{Chars: "abc"}
	rxb := &regexengine.RxBase{}
	text := "abcdef"
	expected := true
	result := rxb.Match(matcher, text)
	if result != expected {
		t.Errorf("Expected match with text '%s' to be %t, got %t", text, expected, result)
	}
}

func TestRxLit_Match(t *testing.T) {
	// MockLitMatcher with non-matching chars
	matcher := &regexengine.RxLit{Chars: "xyz"}
	rxb := &regexengine.RxBase{}
	expected := false
	text := "abcdef"

	result := rxb.Match(matcher, text)
	if result != expected {
		t.Errorf("Expected match with text '%s' to be %t, got %t", text, expected, result)
	}

	// Using NewRegexLit function to create matcher
	matcher = regexengine.NewRegexLit("ab", nil)
	text = "ab"
	expected = true
	result = rxb.Match(matcher, text)
	if result != expected {
		t.Errorf("Expected match with text '%s' to be %t, got %t", text, expected, result)
	}
}

func TestRxStart_MatchThis(t *testing.T) {
	// Test case: start is 0, rest is nil
	rs := regexengine.NewRxStart(nil)
	text := "abcdef"
	start := 0
	expected := 0
	result := rs.MatchThis(text, start)
	if result != expected {
		t.Errorf("Expected match with text '%s' at start %d to return %d, got %d", text, start, expected, result)
	}

	// Test case: start is not 0, corrected expectation
	start = 2
	expected = -1
	result = rs.MatchThis(text, start)
	if result != expected {
		t.Errorf("Expected match with text '%s' at start %d to return %d, got %d", text, start, expected, result)
	}

	// Reset start to 0 for the next case
	start = 0

	// Test case: start is 0, rest is not nil
	rest := regexengine.NewRxStart(nil)
	rs = regexengine.NewRxStart(rest)
	expected = 0
	result = rs.MatchThis(text, start)
	if result != expected {
		t.Errorf("Expected match with text '%s' at start %d to return %d, got %d", text, start, expected, result)
	}
}
