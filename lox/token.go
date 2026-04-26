package main

import "fmt"

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

type Token struct {
	tokenType  TokenType
	lexeme     string
	lineNumber int
	literal    interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("%v %v %v", t.tokenType, t.lexeme, t.literal)
}

func NewToken(tokenType TokenType, lexeme string, lineNumber int, literal interface{}) Token {
	newToken := Token{tokenType: tokenType, lexeme: lexeme, lineNumber: lineNumber, literal: literal}
	return newToken
}
