package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	log.Println("Go-Lox interpreter started.")
	args := os.Args
	if len(args) > 2 {
		log.Fatal("[Error] Correct usage: 'go run main.go [script_name].'")
	} else if len(args) == 2 {
		runFile(args[1])
	} else {
		runPrompt()
	}
}

func runFile(arg string) {
	bytes, err := os.ReadFile(arg)
	if err != nil {
		log.Fatal("[Error] Unable to read file.")
	}
	run(string(bytes))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Interactive mode. Enter Lox code below. Ctrl+D to exit.")
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		run(strings.TrimSpace(line))
	}
}

func run(source string) {
	fmt.Println(source)
}