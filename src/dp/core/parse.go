package dp

import (
	"dp/err"
	"dp/util/algo"
	"fmt"
	"strings"
)

const (
	NOT_FOUND       rune = 0
	FOUND_ENDING    rune = 1
	FOUND_BEGINNING rune = 2
)

const (
	BRANCH_TOKEN       = "->"
	DIMENSION_SEP      = ","
	QUOTE_SMALL_RIGHT  = ')'
	QUOTE_SMALL_LEFT   = '('
	QUOTE_MIDDLE_LEFT  = '['
	QUOTE_MIDDLE_RIGHT = ']'
	QUOTE_LARGE_LEFT   = '{'
	QUOTE_LARGE_RIGHT  = '}'
	QUOTE_SHARP_LEFT   = '<'
	QUOTE_SHARP_RIGTH  = '>'
)

const (
	QUOTE_SMALL  = 0x023300
	QUOTE_MIDDLE = 0x023301
	QUOTE_LARGE  = 0x023302
	QUOTE_SHARP  = 0x023303
)

/// clean the expression
func Clean(str string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(
						str, "and", "&&", -1),
					"or", "||", -1),
				"\n", "", -1),
			" ", "", -1),
		"\t", "", 1)
}

func checkQuote(str string) {
	stack := algo.NewStack(1000)
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case QUOTE_LARGE_LEFT:
			stack.Push(QUOTE_LARGE)
			break
		case QUOTE_MIDDLE_LEFT:
			stack.Push(QUOTE_MIDDLE)
			break
		case QUOTE_SMALL_LEFT:
			stack.Push(QUOTE_SMALL)
			break
		case QUOTE_SHARP_LEFT:
			stack.Push(QUOTE_SHARP)
			break
		case QUOTE_LARGE_RIGHT:
			//
		}
	}
}

func newCondition(str string) string {
	var quoteState = NOT_FOUND
	var endingIndex = -1
	var beginningIndex = -1
	for i := len(str) - 1; i >= 0; i-- {
		switch quoteState {
		case NOT_FOUND:
			if str[i] == QUOTE_SMALL_RIGHT {
				quoteState = FOUND_ENDING
				endingIndex = i
			}
			break
		case FOUND_ENDING:
			if str[i] == QUOTE_SMALL_LEFT {
				quoteState = FOUND_BEGINNING
				beginningIndex = i + 1
				if endingIndex != -1 {
					return str[beginningIndex:endingIndex]
				} else {
					panic("ending quote not found")
				}
			}
			break
		case FOUND_BEGINNING:
			panic("program has been mysteriously exited")
			break
		}
	}
	panic("not valid Expression")
}

func newExpression(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == QUOTE_SMALL_LEFT || str[i] == QUOTE_SMALL_RIGHT {
			return str[:i]
		}
	}
	return str
}

func newBranch(code string) *branch {
	ret := new(branch)
	ret.Conditions = newCondition(code)
	ret.Expression = newExpression(code)
	return ret
}

func newState(str string) *state {
	ret := new(state)
	sep := strings.Index(str, string(QUOTE_MIDDLE_LEFT))
	if sep == -1 {
		err.Raise("main expression error")
	}
	ret.DimExpr = strings.Split(str[sep+1:len(str)-1], DIMENSION_SEP)
	ret.Name = str[:sep]
	return ret
}

func Parse(source string) *dyProInfo {
	ret := newStateEquation(Clean(source))
	fmt.Errorf(err.GetMessages())
	return ret
}

func newStateEquation(source string) *dyProInfo {
	ret := new(dyProInfo)
	ret.Type = "int"
	ret.Detail = *NewImplDetail(10001)
	split := strings.Split(source, BRANCH_TOKEN)
	if len(split) < 2 {
		err.Raise("require branches!")
	} else {
		branches := make([]branch, len(split)-1)
		for index, i := range split[1:] {
			branches[index] = *newBranch(i)
		}
		ret.Branches = branches
		ret.State = *newState(split[0])
	}
	return ret
}
