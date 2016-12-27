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
	BRACE_SMALL_RIGHT  = ')'
	BRACE_SMALL_LEFT   = '('
	BRACE_MIDDLE_LEFT  = '['
	BRACE_MIDDLE_RIGHT = ']'
	BRACE_LARGE_LEFT   = '{'
	BRACE_LARGE_RIGHT  = '}'
	BRACE_SHARP_LEFT   = '<'
	BRACE_SHARP_RIGTH  = '>'
)

const (
	BRACE_SMALL  = 0x00
	BRACE_MIDDLE = 0x01
	BRACE_LARGE  = 0x02
	BRACE_SHARP  = 0x03
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

func checkBrace(str string) {
	stack := algo.NewStack(1000)
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case BRACE_LARGE_LEFT:
			stack.Push(BRACE_LARGE)
			break
		case BRACE_MIDDLE_LEFT:
			stack.Push(BRACE_MIDDLE)
			break
		case BRACE_SMALL_LEFT:
			stack.Push(BRACE_SMALL)
			break
		case BRACE_SHARP_LEFT:
			stack.Push(BRACE_SHARP)
			break
		case BRACE_LARGE_RIGHT:
			if stack.IsEmpty() || stack.Pop() != BRACE_LARGE {
				err.Raise("brace {} doesn't matches!")
			}
			return
		case BRACE_MIDDLE_RIGHT:
			if stack.IsEmpty() || stack.Pop() != BRACE_MIDDLE {
				fmt.Println(stack.Size())
				fmt.Println(stack.Front() == BRACE_MIDDLE)
				err.Raise("brace [] doesn't matches!")
			}
			return
		case BRACE_SMALL_RIGHT:
			if stack.IsEmpty() || stack.Pop() != BRACE_SMALL {
				err.Raise("brace () doesn't matches!")
			}
			return
		case BRACE_SHARP_RIGTH:
			if stack.IsEmpty() || stack.Pop() != BRACE_SHARP {
				err.Raise("brace <> doesn't matches!")
			}
			return
		}
	}
	if !stack.IsEmpty() {
		err.Raise("brace doesn' matches!")
	}
}

func newCondition(str string) string {
	var BRACEState = NOT_FOUND
	var endingIndex = -1
	var beginningIndex = -1
	for i := len(str) - 1; i >= 0; i-- {
		switch BRACEState {
		case NOT_FOUND:
			if str[i] == BRACE_SMALL_RIGHT {
				BRACEState = FOUND_ENDING
				endingIndex = i
			}
			break
		case FOUND_ENDING:
			if str[i] == BRACE_SMALL_LEFT {
				BRACEState = FOUND_BEGINNING
				beginningIndex = i + 1
				if endingIndex != -1 {
					return str[beginningIndex:endingIndex]
				} else {
					err.Raise("ending BRACE not found")
					return "else"
				}
			}
			break
		case FOUND_BEGINNING:
			err.Raise("program has been mysteriously exited")
			return "else"
		}
	}
	err.Raise("require condition!")
	return "else"
}

func newExpression(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == BRACE_SMALL_LEFT || str[i] == BRACE_SMALL_RIGHT {
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
	sep := strings.Index(str, string(BRACE_MIDDLE_LEFT))
	if sep == -1 {
		err.Raise("main expression error")
	}
	ret.DimExpr = strings.Split(str[sep + 1:len(str) - 1], DIMENSION_SEP)
	ret.Name = str[:sep]
	return ret
}

func Parse(source string) *dyProInfo {
	checkBrace(source)
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
		branches := make([]branch, len(split) - 1)
		for index, i := range split[1:] {
			branches[index] = *newBranch(i)
		}
		ret.Branches = branches
		ret.State = *newState(split[0])
	}
	return ret
}
