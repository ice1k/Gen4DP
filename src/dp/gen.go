package dp

import (
	"strings"
)

const (
	NOT_FOUND       rune = 0
	FOUND_ENDING    rune = 1
	FOUND_BEGINNING rune = 2
)

func Clean(str string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(
				str, "\n", "", -1),
			" ", "", -1),
		"\t", "", 1)
}

func GetCondition(str string) string {
	var quoteState = NOT_FOUND
	var endingIndex = -1
	var beginningIndex = -1
	for i := len(str) - 1; i >= 0; i-- {
		switch quoteState {
		case NOT_FOUND:
			if str[i] == ')' {
				quoteState = FOUND_ENDING
				endingIndex = i
			}
			break
		case FOUND_ENDING:
			if str[i] == '(' {
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
	panic("not valid expression")
}

func GetExpression(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == ')' {
			return str[:i]
		}
	}
	return str
}


