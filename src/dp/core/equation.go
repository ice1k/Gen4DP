package dp

/// example: dp[i, j] -> dp[i - 1, j - 1] + 1 (else)
type state struct {
	Name         string /// like: dp
	NameExpr     string /// like: dp[i, j]
	RelationExpr []string
	DimExpr      []string /// Dimension of the dp equ
}

type dyProInfo struct {
	Source   string
	Type     string
	State    state
	Branches []branch
	Detail   implDetail
}

type implDetail struct {
	MaxLen int
}

type branch struct {
	Condition  string
	Expression string
	IsDefault  bool
}

func NewImplDetail(maxLen int) *implDetail {
	ret := new(implDetail)
	ret.MaxLen = maxLen
	return ret
}

/**
package main
const (
	OneDimension = "#include <stdio.h>\n" +
		"int dp[]\n" +
		"int main(const int argc, const char *argv[]) {\n" +
		"\treturn 0;\n" +
		"}\n" +
		"\n"
)
*/
