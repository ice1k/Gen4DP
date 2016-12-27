package dp

import (
	"testing"
	"dp/err"
	"fmt"
)

func TestCheckBrace(t *testing.T) {
	code := "dp[i]] -> dp[[i-1]"
	code2 := "dp[i]] -> dp[[i-1] (i >= 0)"
	res := Parse(code)
	Parse(code2)
	fmt.Println(err.GetMessages())
	fmt.Println(res.GenerateClang(*NewCodeStyle()))
}

func TestCheckBrace2(t *testing.T) {
	code := "dp[1] -> dp[i-1]"
	res := Parse(code)
	fmt.Println(err.GetMessages())
	fmt.Println(res.GenerateClang(*NewCodeStyle()))
}
