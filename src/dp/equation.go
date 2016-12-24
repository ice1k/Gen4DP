package dp

import (
	"strings"
)

type state struct {
	/// dimension of the dp equ
	dimension rune
	/// this is a little bit difficult to explain
	/// document will be available soon
	selfRelations     rune
	externalRelations rune
}

type stateEquation struct {
	stt      state
	branches []branch
}

type equation struct {
}

type branch struct {
	conditions string
	expression string
}

func NewState(source string) *state {
	split := strings.Split(source, "=")
	if len(split) < 2 {
		return nil
	} else {
		branches := make([]branch, len(split) - 1)
		for index, i := range split[2:] {
			branches[index].conditions = GetCondition(Clean(i))
			branches[index].expression = GetExpression(Clean(i))
		}
	}
	return nil
}

//const (
//	OneDimention = "#include <stdio.h>\n" +
//		"int dp[]\n" +
//		"int main(const int argc, const char *argv[]) {\n" +
//		"\treturn 0;\n" +
//		"}\n" +
//		"\n"
//)
