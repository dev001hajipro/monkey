package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("こんにちは %s!, This is the 🐵Monkey programming langauge!\n", user.Username)
	fmt.Printf("コマンドどうぞ\n")
}