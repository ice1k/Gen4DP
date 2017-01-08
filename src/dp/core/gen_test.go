package core

import (
	"dp/err"
	"dp/msg"
	"fmt"
	"testing"
)

func TestCorrect(t *testing.T) {
	code := `
dp[i] -> dp[i - 1] + dp[i - 2] (i >= 2)
      -> 1 (i == 1 or i == 2)
      -> 0 (else)
	`
	res := Parse(code)
	fmt.Println(msg.GetErrors())
	fmt.Println(res.GenerateClang(*NewCodeStyle()))
}

func TestExample(t *testing.T) {
	code := `
dp[i,j] -> dp[i - 1, j] * i (i >= 1)
      -> 1 (i == 0)`
	fmt.Println(Parse(code).GenerateClang(*NewCodeStyle()))
}
