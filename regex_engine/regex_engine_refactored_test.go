package regexengine_test

import (
	regexengine "softwaredesign/regex_engine"
	"testing"
)

func TestRxBase_Match(t *testing.T) {

	// MockLitMatcher with matching chars
	matcher := &regexengine.RxLit{"abc"}
	rxb := &regexengine.RxBase{}
	text := "abcdef"
	expected := true
	result := rxb.Match(matcher, text)
	if result != expected {
		t.Errorf("Expected match with text '%s' to be %t, got %t", text, expected, result)
	}

	// MockLitMatcher with non-matching chars
	matcher = &regexengine.RxLit{Chars: "xyz"}
	expected = false
	result = rxb.Match(matcher, text)
	if result != expected {
		t.Errorf("Expected match with text '%s' to be %t, got %t", text, expected, result)
	}
}
