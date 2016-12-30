package dp

import (
	"dp/err"
	"dp/util/sb"
	"fmt"
	"os"
	"os/exec"
	"bufio"
	"strings"
)

const (
	COMMAND_BEGIN = ">> "
	COMMAND_INFIX = "\t | "
)

func StartRepl() {
	fmt.Println("Gen4DP v1.0, open source under Apache 2.0 license.")
	fmt.Print(">> ")
	buffer := sb.NewStringBuffer()
	bio := bufio.NewReader(os.Stdin)
	var s []byte
	for true {
		s, _, _ = bio.ReadLine()
		//fmt.Println(string(s))
		switch  strings.Trim(string(s), " \n\t") {
		case "end":
			in := buffer.ToString()
			fmt.Println(in)
			fmt.Println(Parse(in).GenerateClang(*NewCodeStyle()))
			fmt.Println(err.GetErrors())
			buffer.Clear()
			err.Clear()
			fmt.Print(COMMAND_BEGIN)
		case "cls":
			exec.Command("cls")
			fmt.Print(COMMAND_BEGIN)
		case "exit":
			err.Clear()
			buffer.Clear()
			fmt.Println("Have a nice day. (^_^)")
			return
		default:
			buffer.Append(string(s))
			fmt.Print(COMMAND_INFIX)
		}
	}
}
