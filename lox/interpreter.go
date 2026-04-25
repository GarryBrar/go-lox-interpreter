package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Interpreter struct{
	hadError bool
}

func (i *Interpreter) runFile(arg string) {
	bytes, err := os.ReadFile(arg)
	if err != nil {
		log.Fatal("[Error] Unable to read file.")
	}
	i.run(string(bytes))
}

func (i *Interpreter) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Interactive mode. Enter Lox code below. Ctrl+D to exit.")
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		i.run(strings.TrimSpace(line))
	}
}

func (i *Interpreter) run(source string) {
	fmt.Println(source)
}