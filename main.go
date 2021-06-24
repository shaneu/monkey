package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/shaneu/monkey/pkg/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language.\n", user.Username)
	fmt.Println("type in your commands")
	repl.Start(os.Stdin, os.Stdout)
}
