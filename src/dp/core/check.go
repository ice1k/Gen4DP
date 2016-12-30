package dp

import (
	"dp/err"
	"dp/util/algo"
	"strings"
)

const (
	BRACE_SMALL_RIGHT  = ')'
	BRACE_SMALL_LEFT   = '('
	BRACE_MIDDLE_LEFT  = '['
	BRACE_MIDDLE_RIGHT = ']'
	BRACE_LARGE_LEFT   = '{'
	BRACE_LARGE_RIGHT  = '}'
)

const (
	BRACE_SMALL  = 0x000
	BRACE_MIDDLE = 0x001
	BRACE_LARGE  = 0x002
)

/// get C++ preserved words
func getPreservedWords() []string {
	return []string{
		"and",
		"asm",
		"auto",
		"and_eq",
		"bad_cast",
		"bad_typeid",
		"bitand",
		"bitor",
		"or_eq",
		"bool",
		"break",
		"case",
		"catch",
		"char",
		"class",
		"const",
		"const_cast",
		"char16_t",
		"char32_t",
		"__restrict__",
		"__cdecl",
		"static_assert",
		"continue",
		"default",
		"delete",
		"do",
		"double",
		"dynamic_cast",
		"else",
		"enum",
		"except",
		"explicit",
		"extern",
		"false",
		"finally",
		"float",
		"for",
		"friend",
		"goto",
		"if",
		"inline",
		"int",
		"long",
		"mutable",
		"namespace",
		"new",
		"operator",
		"or",
		"private",
		"protected",
		"public",
		"reinterpret_cast",
		"return",
		"short",
		"signed",
		"sizeof",
		"static",
		"static_cast",
		"struct",
		"switch",
		"template",
		"this",
		"throw",
		"true",
		"try",
		"type_info",
		"typedef",
		"typeid",
		"typename",
		"union",
		"unsigned",
		"using",
		"virtual",
		"void",
		"volatile",
		"wchar_t",
		"while",
		"xor",
		"xor_eq",
		"register"}
}

func getBraceLeft(mark int) byte {
	switch mark {
	case BRACE_SMALL:
		return BRACE_SMALL_LEFT
	case BRACE_LARGE:
		return BRACE_LARGE_LEFT
	case BRACE_MIDDLE:
		return BRACE_MIDDLE_LEFT
	}
	panic("Mark not found")
}

func getBraceRight(mark int) byte {
	switch mark {
	case BRACE_SMALL:
		return BRACE_SMALL_RIGHT
	case BRACE_LARGE:
		return BRACE_LARGE_RIGHT
	case BRACE_MIDDLE:
		return BRACE_MIDDLE_RIGHT
	}
	panic("Mark not found")
}

func getMark(left byte) int {
	switch left {
	case BRACE_SMALL_LEFT:
		return BRACE_SMALL
	case BRACE_LARGE_LEFT:
		return BRACE_LARGE
	case BRACE_MIDDLE_LEFT:
		return BRACE_MIDDLE
	}
	return -1
}

func isLeftBrace(i byte) bool {
	return i == BRACE_LARGE_LEFT ||
		i == BRACE_MIDDLE_LEFT ||
		i == BRACE_SMALL_LEFT
}

func isRightBrace(i byte) bool {
	return i == BRACE_LARGE_RIGHT ||
		i == BRACE_MIDDLE_RIGHT ||
		i == BRACE_SMALL_RIGHT
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isValid(b byte) bool {
	return isAlpha(b) || b == '$' || b == '_'
}

func isOperator(b byte) bool {
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

func checkBrace(str string) {
	stack := algo.NewStack(100)
	line := 1
	row := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '\n' {
			row++
		} else {
			line++
			row = 0
		}
		if isLeftBrace(str[i]) {
			stack.Push(getMark(str[i]))
		} else if isRightBrace(str[i]) {
			if stack.IsEmpty() {
				err.RaiseFormat("Unexpected '%c'.", str[i])
			} else if getBraceRight(stack.Top()) != str[i] {
				err.RaiseFormat(
					"Brace doesn't match, expected '%c', found '%c' at line %d, row %d.",
					getBraceRight(stack.Top()),
					str[i],
					line,
					row)
			}
			stack.Pop()
		}
	}
	if !stack.IsEmpty() {
		err.Raise("Expected ending brace.")
	}
}

func runSymbolCheck(name string) bool {
	isNum := false
	lastOneIsOperator := true
	for j := 0; j < len(name); j++ {
		if lastOneIsOperator {
			lastOneIsOperator = false
			if !isValid(name[j]) {
				isNum = true
			}
		}
		if isOperator(name[j]) {
			isNum = false
			lastOneIsOperator = true
			continue
		}
		if !isDigit(name[j]) {
			if isNum || !isAlpha(name[j]) {
				return false
			}
		}
	}
	return true
}

func runExprDimCheck(expr string, dim int) bool {
	inBrace := false
	lastBraceIndex := -1
	for j := 0; j < len(expr); j++ {
		if !inBrace && expr[j] == '[' {
			inBrace = true
			lastBraceIndex = j
		}
		if inBrace && expr[j] == ']' {
			inBrace = false
			foundDim := strings.Count(expr[lastBraceIndex:j], ",") + 1
			if foundDim != dim {
				err.RaiseFormat(
					"Require %d indexes, found %d.",
					dim,
					foundDim)
				return false
			}
		}
	}
	return true
}

func checkSymbol(info *dyProInfo) {
	for _, i := range info.State.DimExpr {
		if !runSymbolCheck(i) {
			err.RaiseFormat("Invalid name in expression: %s.", i)
		}
	}
}

func checkDimension(info *dyProInfo) {
	for i := 0; i < len(info.State.DimExpr); i++ {
		for j := i + 1; j < len(info.State.DimExpr); j++ {
			if info.State.DimExpr[i] == info.State.DimExpr[j] {
				err.RaiseFormat("Dimension index name redefined: %s.",
					info.State.DimExpr[i])
			}
		}
	}
	for _, i := range info.Branches {
		runExprDimCheck(i.Expression, len(info.State.DimExpr))
	}
}
