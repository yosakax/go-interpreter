package main

import (
	"fmt"
	"os"
	"os/user"

	"yosaka/go-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %s! This is the Monkey programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
