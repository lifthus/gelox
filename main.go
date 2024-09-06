package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/lifthus/gelox/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Gelox engine!\n", user.Username)
	fmt.Printf("Type something!\n")
	repl.Start(os.Stdin, os.Stdout)
}
