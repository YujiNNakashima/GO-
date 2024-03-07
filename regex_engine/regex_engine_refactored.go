package regexengine

type RxBase struct {
}

func (rxb *RxBase) Match(m Matcher, text string) bool {
	for i := 0; i < len(text); i++ {
		if m.MatchThis(text, i) {
			return true
		}
	}
	return false
}

type Matcher interface {
	MatchThis(text string, start int) bool
}

type RxLit struct {
	Chars string
}

func (m *RxLit) MatchThis(text string, start int) bool {
	nextIdx := start + len(m.Chars)
	if nextIdx > len(text) {
		return false
	}
	substr := text[start:nextIdx]
	return substr == m.Chars
}
