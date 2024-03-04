package regexengine

func Match(pattern, text string) bool {
	if len(pattern) > 0 && pattern[0] == '^' {
		return matchFromHere([]byte(pattern), 1, text, 0)
	}

	for idxText := 0; idxText < len(text); idxText++ {
		if matchFromHere([]byte(pattern), 0, text, idxText) {
			return true
		}
	}

	return false
}

func matchFromHere(pattern []byte, idxPattern int, text string, idxText int) bool {

	if idxPattern == len(pattern) {
		return true
	}

	// pattern '$'
	if idxPattern == len(pattern)-1 && pattern[idxPattern] == '$' && idxText == len(text) {
		return true
	}

	// pattern '*'
	if idxPattern < len(pattern)-1 && pattern[idxPattern+1] == '*' {
		for idxText < len(text) && (pattern[idxPattern] == '.' || text[idxText] == pattern[idxPattern]) {
			idxText++
		}
		return matchFromHere(pattern, idxPattern+2, text, idxText)
	}

	// pattern .
	if idxText < len(text) && (pattern[idxPattern] == '.' || pattern[idxPattern] == text[idxText]) {
		return matchFromHere(pattern, idxPattern+1, text, idxText+1)
	}

	return false
}
