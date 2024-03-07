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

func TestRxEnd_MatchThis(t *testing.T) {
	// Test case 1: Match at the end of the string
	end := regexengine.NewRxEnd(nil)
	text := "hello"
	if got := end.MatchThis(text, len(text)); got != len(text) {
		t.Errorf("RxEnd.MatchThis() = %v, want %v", got, len(text))
	}

	// Test case 2: Not matching if start is not at the end
	if got := end.MatchThis(text, 0); got != -1 {
		t.Errorf("RxEnd.MatchThis() = %v, want %v", got, -1)
	}

	// Test case 3: Chaining with another matcher
	mockMatcher := &MockMatcher{expectedStart: len(text)}
	endWithRest := regexengine.NewRxEnd(mockMatcher)
	if got := endWithRest.MatchThis(text, len(text)); got != mockMatcher.expectedStart {
		t.Errorf("RxEnd.MatchThis() with rest = %v, want %v", got, mockMatcher.expectedStart)
	}
}

type MockMatcher struct {
	expectedStart int
}

func (mm *MockMatcher) MatchThis(text string, start int) int {
	return mm.expectedStart
}

func TestRxAlt_MatchThis(t *testing.T) {
	leftMock := &MockMatcher{expectedStart: 3}  // Assume it matches and moves the start to 3
	rightMock := &MockMatcher{expectedStart: 5} // Assume it matches and moves the start to 5
	text := "hello world"

	// Test case 1: Left matcher succeeds
	altLeft := regexengine.NewRxAlt(leftMock, nil, nil)
	if got := altLeft.MatchThis(text, 0); got != 3 {
		t.Errorf("RxAlt.MatchThis() left = %v, want %v", got, 3)
	}

	// Test case 2: Right matcher succeeds
	altRight := regexengine.NewRxAlt(nil, rightMock, nil)
	if got := altRight.MatchThis(text, 0); got != 5 {
		t.Errorf("RxAlt.MatchThis() right = %v, want %v", got, 5)
	}

	// Test case 3: Chaining with another matcher
	restMock := &MockMatcher{expectedStart: len(text)} // Assume rest matches to the end
	altChain := regexengine.NewRxAlt(leftMock, rightMock, restMock)
	if got := altChain.MatchThis(text, 0); got != len(text) {
		t.Errorf("RxAlt.MatchThis() chain = %v, want %v", got, len(text))
	}

	// Test case 4: Neither pattern matches
	altNone := regexengine.NewRxAlt(&MockMatcher{expectedStart: -1}, &MockMatcher{expectedStart: -1}, nil)
	if got := altNone.MatchThis(text, 0); got != -1 {
		t.Errorf("RxAlt.MatchThis() none = %v, want %v", got, -1)
	}
}
