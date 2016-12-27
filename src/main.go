package main

import (
	"./dp/core"
	"fmt"
)

func main() {
	var equ = `
dp[i] -> dp[i - 1] + dp[i - 2] + dp[i - 3](i >= 2)
      -> 2 (i == 2 or i == 3)
      -> 1 (i == 1)
      -> 0 (else)`
	fmt.Println(dp.Parse(equ).GenerateClang(*dp.NewCodeStyle()))
}
