package core

import (
	"dp/msg"
	"dp/util"
	"dp/util/algo"
	"strings"
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
		return util.BRACE_SMALL_LEFT
	case BRACE_LARGE:
		return util.BRACE_LARGE_LEFT
	case BRACE_MIDDLE:
		return util.BRACE_MIDDLE_LEFT
	}
	panic("Mark not found")
}

func getBraceRight(mark int) byte {
	switch mark {
	case BRACE_SMALL:
		return util.BRACE_SMALL_RIGHT
	case BRACE_LARGE:
		return util.BRACE_LARGE_RIGHT
	case BRACE_MIDDLE:
		return util.BRACE_MIDDLE_RIGHT
	}
	panic("Mark not found")
}

func getMark(left byte) int {
	switch left {
	case util.BRACE_SMALL_LEFT:
		return BRACE_SMALL
	case util.BRACE_LARGE_LEFT:
		return BRACE_LARGE
	case util.BRACE_MIDDLE_LEFT:
		return BRACE_MIDDLE
	}
	return -1
}

func checkBrace(str string, sig *chan byte) {
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
		if util.IsLeftBrace(str[i]) {
			stack.Push(getMark(str[i]))
		} else if util.IsRightBrace(str[i]) {
			if stack.IsEmpty() {
				msg.RaiseFormat("Unexpected '%c'.", str[i])
			} else if getBraceRight(stack.Top()) != str[i] {
				msg.RaiseFormat(
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
		msg.Raise("Expected ending brace.")
	}
	<-*sig
}

func runSymbolCheck(name string) bool {
	isNum := false
	lastOneIsOperator := true
	for j := 0; j < len(name); j++ {
		if lastOneIsOperator {
			lastOneIsOperator = false
			if !util.IsValid(name[j]) {
				isNum = true
			}
		}
		if util.IsOperator(name[j]) {
			isNum = false
			lastOneIsOperator = true
			continue
		}
		if !util.IsDigit(name[j]) {
			if isNum || !util.IsAlpha(name[j]) {
				return false
			}
		}
	}
	return true
}

func isPreservedWord(str string) bool {
	for _, i := range getPreservedWords() {
		if i == str {
			return true
		}
	}
	return false
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
				msg.RaiseFormat(
					"Require %d indexes, found %d.",
					dim,
					foundDim)
				return false
			}
		}
	}
	return true
}

func checkSymbol(info *dyProInfo, sig *chan byte) {
	chk := func(i string) {
		if isPreservedWord(i) {
			msg.RaiseFormat("C++ preserved word: %s.", i)
			return
		}
		if !runSymbolCheck(i) {
			msg.RaiseFormat("Invalid name in expression: %s.", i)
		}
	}
	for _, i := range info.State.DimExpr {
		chk(i)
	}
	chk(info.State.Name)
	<-*sig
}

func checkDimension(info *dyProInfo, sig *chan byte) {
	for i := 0; i < len(info.State.DimExpr); i++ {
		if isPreservedWord(info.State.DimExpr[i]) {
			msg.RaiseFormat("C++ preserved word: %s.",
				info.State.DimExpr[i])
			continue
		}
		for j := i + 1; j < len(info.State.DimExpr); j++ {
			if info.State.DimExpr[i] == info.State.DimExpr[j] {
				msg.RaiseFormat("Dimension index name redefined: %s.",
					info.State.DimExpr[i])
			}
		}
	}
	for _, i := range info.Branches {
		runExprDimCheck(i.Expression, len(info.State.DimExpr))
	}
	<-*sig
}
