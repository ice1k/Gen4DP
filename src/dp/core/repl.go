package core

import (
	"bufio"
	"dp/msg"
	"dp/util"
	"dp/util/sb"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	COMMAND_BEGIN = ">> "
	COMMAND_INFIX = "\t | "
)

func StartRepl() {
	fmt.Println("Gen4DP v1.0, open source under GNU General Public License v3.0.")
	fmt.Print(">> ")
	buffer := sb.NewStringBuffer()
	bio := bufio.NewReader(os.Stdin)
	var s []byte
	for true {
		s, _, _ = bio.ReadLine()
		cmd := strings.Trim(string(s), " \n\t")
		switch cmd {
		case "end":
			in := buffer.ToString()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println(in)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println(Parse(in).GenerateClang(*NewCodeStyle()))
			fmt.Println(msg.GetErrors())
			buffer.Clear()
			fmt.Print(COMMAND_BEGIN)
		case "cls":
			exec.Command("cls").Run()
			fmt.Print(COMMAND_BEGIN)
		case "err":
			if msg.HasError() {
				fmt.Println(msg.GetErrors())
			} else {
				fmt.Println("No errors.")
			}
			fmt.Print(COMMAND_BEGIN)
		case "help":
			fmt.Println(`help:

Equation is something like:
=== === === === === ===
a[i] -> a[i - 1] + i (i >= 1)
     -> 0 (else)
     end
=== === === === === ===

command format: [command]
command list:
	cls: clear screen
	help: see the help document
	end: tell the repl to stop reading an equation
	exit: exit the repl`)
			fmt.Print(COMMAND_BEGIN)
		case "exit":
			msg.Clear()
			buffer.Clear()
			fmt.Println("Have a nice day. (^_^)")
			return
		default:
			if util.AreAlpha(cmd) {
				fmt.Printf("Command \"%s\" not found. Need help? Type: \"help\"\n", cmd)
				fmt.Print(COMMAND_BEGIN)
			} else {
				msg.Clear()
				buffer.Append(string(s))
				fmt.Print(COMMAND_INFIX)
			}
		}
	}
}
