package dp

import (
	"fmt"
	"dp"
)

func ParseTest() {
	var str = dp.Clean("a[i, j] = a[i - 1, j - 1] + 1 ( 2333 )")
	fmt.Println(dp.GetCondition(str))
	fmt.Println(dp.GetExpression(str))
	var equ = `
dp[i] = dp[i - 1] + dp[i - 2] (i >= 2)
	= 1 (i == 1 || i == 2)
	= 0 (else)`
	state := dp.NewState(equ)
	fmt.Println(state)
}
