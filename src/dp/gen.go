package dp

func GetCondition(str string) string {
	const (
		NOT_FOUND       rune = 0
		FOUND_ENDING    rune = 1
		FOUND_BEGINNING rune = 2
	)
	var quoteState = NOT_FOUND
	var endingIndex = -1
	var beginningIndex = -1
	for i := len(str) - 1; i >= 0; i-- {
		switch quoteState {
		case NOT_FOUND:
			if str[i] == ')' {
				quoteState = FOUND_ENDING
				endingIndex = i - 1
			}
			break
		case FOUND_ENDING:
			if str[i] == '(' {
				quoteState = FOUND_BEGINNING
				beginningIndex = i
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
	return "" // TODO
}
