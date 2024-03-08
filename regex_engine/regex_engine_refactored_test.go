package regexengine_test

import (
	regexengine "softwaredesign/regex_engine"
	"testing"
)

type MockMatcher struct {
	matchFunc func(text string, start int) int
}

func (mm *MockMatcher) MatchThis(text string, start int) int {
	if mm.matchFunc != nil {
		return mm.matchFunc(text, start)
	}
	return -1
}

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
	mockMatcher := &MockMatcher{matchFunc: func(text string, start int) int {
		// Assuming rest matches to the end of the text
		return len(text)
	}}
	endWithRest := regexengine.NewRxEnd(mockMatcher)
	expectedMatchPosition := len(text)

	// Use the expectedMatchPosition directly in your assertion
	if got := endWithRest.MatchThis(text, len(text)); got != expectedMatchPosition {
		t.Errorf("RxEnd.MatchThis() with rest = %v, want %v", got, expectedMatchPosition)
	}
}

func TestRxAlt_MatchThis(t *testing.T) {
	leftMock := &MockMatcher{
		matchFunc: func(text string, start int) int {
			// Assuming it matches and moves the start to 5
			return 3
		},
	}
	rightMock := &MockMatcher{
		matchFunc: func(text string, start int) int {
			// Assuming it matches and moves the start to 5
			return 5
		},
	}
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
	restMock := &MockMatcher{
		matchFunc: func(text string, start int) int {
			// Assuming rest matches to the end of the text
			return len(text)
		},
	}
	altChain := regexengine.NewRxAlt(leftMock, rightMock, restMock)
	if got := altChain.MatchThis(text, 0); got != len(text) {
		t.Errorf("RxAlt.MatchThis() chain = %v, want %v", got, len(text))
	}

	// Test case 4: Neither pattern matches
	altNone := regexengine.NewRxAlt(&MockMatcher{
		matchFunc: func(text string, start int) int {
			// This matcher always fails to match, returning -1
			return -1
		},
	}, &MockMatcher{
		matchFunc: func(text string, start int) int {
			// This matcher also always fails to match, returning -1
			return -1
		},
	}, nil) // Corrected by removing the extra comma

	if got := altNone.MatchThis(text, 0); got != -1 {
		t.Errorf("RxAlt.MatchThis() none = %v, want %v", got, -1)
	}
}

func TestRxAny_MatchThis(t *testing.T) {
	singleAMatcher := &MockMatcher{
		matchFunc: func(text string, start int) int {
			if start < len(text) && text[start] == 'a' {
				return start + 1
			}
			return -1
		},
	}

	// Test case 1: Match multiple 'a's and return the new start position
	text := "aaab"
	anyA := regexengine.NewRxAny(singleAMatcher, nil)
	if got := anyA.MatchThis(text, 0); got != 3 {
		t.Errorf("RxAny.MatchThis() = %v, want %v", got, 3)
	}

	// Test case 2: Chain to another matcher after matching multiple 'a's
	bMatcher := regexengine.NewRegexLit("b", nil)
	anyAWithRest := regexengine.NewRxAny(singleAMatcher, bMatcher)
	if got := anyAWithRest.MatchThis(text, 0); got != len(text) {
		t.Errorf("RxAny.MatchThis() with rest = %v, want %v", got, len(text))
	}
}
