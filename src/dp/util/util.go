package util

const (
	BRACE_SMALL_RIGHT  = ')'
	BRACE_SMALL_LEFT   = '('
	BRACE_MIDDLE_LEFT  = '['
	BRACE_MIDDLE_RIGHT = ']'
	BRACE_LARGE_LEFT   = '{'
	BRACE_LARGE_RIGHT  = '}'
)

func IsLeftBrace(i byte) bool {
	return i == BRACE_LARGE_LEFT ||
		i == BRACE_MIDDLE_LEFT ||
		i == BRACE_SMALL_LEFT
}

func IsRightBrace(i byte) bool {
	return i == BRACE_LARGE_RIGHT ||
		i == BRACE_MIDDLE_RIGHT ||
		i == BRACE_SMALL_RIGHT
}

func IsAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func IsValid(b byte) bool {
	return IsAlpha(b) || b == '$' || b == '_'
}

func IsOperator(b byte) bool {
	return b == '+' ||
		b == '-' ||
		b == '/' ||
		b == '|' ||
		b == '*' ||
		b == '(' ||
		b == ')' ||
		b == '%' ||
		b == '&' ||
		b == '!' ||
		b == '~' ||
		b == '^'
}

/// allows alpha an space and '\n'
func AreAlpha(b string) bool {
	for i := 0; i < len(b); i++ {
		if !(IsAlpha(b[i]) ||
			b[i] == ' ' ||
			b[i] == '\n') {
			return false
		}
	}
	return true
}
