package dp

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var equ = `
dp[i, j, k] -> dp[i - 1] + dp[i - 2] (i >= 2)
      -> 1 (i == 1 or i == 2)
      -> 0 (else)`
	equation := Parse(equ)
	fmt.Println("Name:", equation.State.Name)
	fmt.Println("Dims:", len(equation.State.DimExpr))
	for i := 0; i < len(equation.Branches); i++ {
		fmt.Print("In branch ", i, ":\n\texpression:\n\t\t", equation.Branches[i].Expression)
		fmt.Print("\n\tcondition:\n\t\t", equation.Branches[i].Conditions)
		fmt.Println()
	}
	fmt.Printf("max len = %d\n", equation.Detail.MaxLen)
	style := NewCodeStyle()
	fmt.Println(equation.GenerateClang(*style))
}

func TestParse2(t *testing.T) {
	var equ = `
dp[i] -> dp[i - 1] + dp[i - 2] + dp[i - 3](i >= 2)
      -> 2 (i == 2 or i == 3)
      -> 1 (i == 1)
      -> 0 (else)`
	fmt.Println(Parse(equ).GenerateClang(*NewCodeStyle()))
}
