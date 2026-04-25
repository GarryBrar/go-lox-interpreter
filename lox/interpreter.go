package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Interpreter struct {
	hadError bool
}

func (i *Interpreter) runFile(arg string) {
	bytes, err := os.ReadFile(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] Unable to read file.")
		os.Exit(66)
	}
	i.run(string(bytes))
	if i.hadError {
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
		i.hadError = false
	}
}

func (i *Interpreter) run(source string) {
	fmt.Println(source)
}

func (i *Interpreter) error(line int, message string) {
	i.report(line, "", message)
}

func (i *Interpreter) report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	i.hadError = true
}
