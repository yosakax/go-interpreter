package main

import (
	"fmt"
	"os"
	"os/user"

	"yosaka/go-interpreter/repl"
	"yosaka/go-interpreter/runner"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 0 {
		fmt.Printf("hello %s! This is the Monkey programming language!\n", user.Username)

		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)

	} else {
		fmt.Println("please input filename!")
		runner.Run("./example/01.mk", os.Stdout)
	}
}
