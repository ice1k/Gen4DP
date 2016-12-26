package dp

type state struct {
	Name    string
	DimExpr []string /// Dimension of the dp equ
	//SelfRelationExpr     []string
	//ExternalRelationExpr []string
}

type dyProInfo struct {
	State    state
	Type     string
	Branches []branch
	Detail   implDetail
}

type implDetail struct {
	MaxLen int
}

type branch struct {
	Conditions string
	Expression string
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
