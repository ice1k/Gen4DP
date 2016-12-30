package dp

import (
	"dp/err"
	"dp/util/sb"
	"fmt"
	"strings"
)

func StartRepl() {
	fmt.Println("Gen4DP v1.0, open source under Apache 2.0 license.")
	fmt.Print(">> ")
	sb := sb.NewStringBuffer()
	var s string
	for true {
		fmt.Scanln(&s)
		fmt.Println(s)
		if strings.Trim(s, " \n\t") != "end" {
			sb.Append(s)
			fmt.Print(">> ")
		} else {
			in := sb.ToString()
			fmt.Println(in)
			fmt.Println(Parse(in).GenerateClang(*NewCodeStyle()))
			fmt.Println(err.GetErrors())
			err.Clear()
			break
		}
	}
}
