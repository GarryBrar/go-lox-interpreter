package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go-Lox interpreter started.")

	args := os.Args
	interpreter := NewInterpreter()

	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "[Error] Correct usage: 'go run . [script_name].'")
		os.Exit(64)
	} else if len(args) == 2 {
		interpreter.runFile(args[1])
	} else {
		interpreter.runPrompt()
	}
}
