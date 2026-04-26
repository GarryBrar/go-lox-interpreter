package main

import (
	"fmt"
	"os"
)

type ErrorReporter interface {
	error(line int, message string)
}

type ErrorHandler struct {
	hadError bool
}

func (e *ErrorHandler) error(line int, message string) {
	e.report(line, "", message)
}

func (e *ErrorHandler) report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	e.hadError = true
}
