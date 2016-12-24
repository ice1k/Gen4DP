package dp

import (
	"fmt"
	"dp"
)

func ParseTest() {
	var str = dp.Clean("a[i, j] = a[i - 1, j - 1] + 1 ( 2333 )")
	fmt.Println(dp.GetCondition(str))
	fmt.Println(dp.GetExpression(str))
}
