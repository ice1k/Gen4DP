package dp

import (
	"dp/err"
	"fmt"
	"testing"
)

func TestCheckBrace(t *testing.T) {
	code := "dp[i) -> dp[[i-1] (i >= 0)"
	res := Parse(code)
	fmt.Println(err.GetErrors())
	fmt.Println(res.GenerateClang(*NewCodeStyle()))
}

func TestCheckCondition(t *testing.T) {
	code := "dp[1] -> dp[i - 1]"
	res := Parse(code)
	fmt.Println(err.GetErrors())
	fmt.Println(res.GenerateClang(*NewCodeStyle()))
}

func TestCheckName(t *testing.T) {
	code := "dp[233i - 1] -> dp[boyi - 1] (else)"
	res := Parse(code)
	fmt.Println(err.GetErrors())
	fmt.Println(res.GenerateClang(*NewCodeStyle()))
}

func TestCheckBranch(t *testing.T) {
	code := "dp[idx, index, idx] -> dp[idx] (else)"
	res := Parse(code)
	fmt.Println(err.GetErrors())
	fmt.Printf(res.GenerateClang(*NewCodeStyle()))
}
