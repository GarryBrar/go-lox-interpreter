package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Interpreter struct {
	errors *ErrorHandler
}

func NewInterpreter() *Interpreter {
	return &Interpreter{errors: &ErrorHandler{}}
}

func (i *Interpreter) runFile(arg string) {
	bytes, err := os.ReadFile(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Unable to read file.")
		os.Exit(66)
	}
	i.run(string(bytes))
	if i.errors.hadError {
		os.Exit(65)
	}
}

func (i *Interpreter) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interactive mode. Enter Lox code below. Ctrl+D to exit.")
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nExiting interactive mode.")
			break
		}
		i.run(strings.TrimSpace(line))
		i.errors.hadError = false
	}
}

func (i *Interpreter) run(source string) {
	scanner := NewScanner(source, i.errors)
	for _, token := range scanner.ScanTokens() {
		fmt.Println(token)
	}
}
