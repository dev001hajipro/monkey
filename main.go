package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/dev001hajipro/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("こんにちは %s!, This is the 🐵Monkey programming langauge!\n", user.Username)
	fmt.Printf("コマンドどうぞ (exit: Ctrl+c)\n")

	repl.Start(os.Stdin, os.Stdout)
}
