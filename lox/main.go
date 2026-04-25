package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Go-Lox interpreter started.")

	args := os.Args
	interpreter := &Interpreter{}
	
	if len(args) > 2 {
		log.Fatal("[Error] Correct usage: 'go run main.go [script_name].'")
	} else if len(args) == 2 {
		interpreter.runFile(args[1])
	} else {
		interpreter.runPrompt()
	}
}
