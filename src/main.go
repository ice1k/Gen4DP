package main

import (
	"dp/core"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		for i := 1; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "-f":
			}
		}
	} else {
		core.StartRepl()
	}
}
