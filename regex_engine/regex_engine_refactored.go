package regexengine

type RxBase struct {
}

func (rxb *RxBase) Match(m Matcher, text string) bool {
	for i := 0; i < len(text); i++ {
		if m.MatchThis(text, i) != -1 {
			return true
		}
	}
	return false
}

type Matcher interface {
	MatchThis(text string, start int) int
}

type RxLit struct {
	Chars string
	Rest  *RxLit
}

func (rl *RxLit) MatchThis(text string, start int) int {
	nextIdx := start + len(rl.Chars)
	if nextIdx > len(text) {
		return -1
	}
	substr := text[start:nextIdx]
	if substr != rl.Chars {
		return -1
	}
	if rl.Rest == nil {
		return nextIdx
	}
	return rl.Rest.MatchThis(text, nextIdx)
}

func NewRegexLit(chars string, rest *RxLit) *RxLit {
	return &RxLit{
		Chars: chars,
		Rest:  rest,
	}
}

type RxStart struct {
	Rest *RxStart
}

func (rs *RxStart) MatchThis(text string, start int) int {
	if start != 0 {
		return -1
	}
	if rs.Rest == nil {
		return 0
	}
	return rs.Rest.MatchThis(text, start)
}

func NewRxStart(rest *RxStart) *RxStart {
	return &RxStart{
		Rest: rest,
	}
}

type RxEnd struct {
	Rest Matcher
}

func (re *RxEnd) MatchThis(text string, start int) int {
	if start != len(text) {
		return -1
	}
	if re.Rest == nil {
		return len(text)
	}
	return re.Rest.MatchThis(text, start)
}

func NewRxEnd(rest Matcher) *RxEnd {
	return &RxEnd{
		Rest: rest,
	}
}

type RxAlt struct {
	Left, Right, Rest Matcher
}

func (ra *RxAlt) MatchThis(text string, start int) int {
	patterns := []Matcher{ra.Left, ra.Right}
	for _, pat := range patterns {
		if pat == nil {
			continue
		}
		afterPat := pat.MatchThis(text, start)
		if afterPat != -1 {
			if ra.Rest == nil {
				return afterPat
			}
			afterRest := ra.Rest.MatchThis(text, afterPat)
			if afterRest != -1 {
				return afterRest
			}
		}
	}
	return -1
}

func NewRxAlt(left, right, rest Matcher) *RxAlt {
	return &RxAlt{
		Left:  left,
		Right: right,
		Rest:  rest,
	}
}
