package main

import (
	"fmt"
	"os"

	"monkey.lang/repl"
)

func main() {
	fmt.Printf("Monkey.lang REPL\n")
	fmt.Printf("Type your commands:\n")

	repl.Start(os.Stdin, os.Stdout)
}
