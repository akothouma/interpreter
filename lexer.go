package main

import "strings"

type Lexer struct{}

func (l *Lexer) Tokenize(input string) []string {
	return strings.Fields(input)
}