package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Store variables
var variables = make(map[string]int)

// Tokenize input into commands
func tokenize(input string) []string {
	return strings.Fields(input)
}

// Interpret and execute commands
func interpret(tokens []string) {
	if len(tokens) == 0 {
		return
	}

	switch tokens[0] {
	case "print":
		if len(tokens) > 1 {
			if val, ok := variables[tokens[1]]; ok {
				fmt.Println(val)
			} else {
				fmt.Println(tokens[1])
			}
		}
	case "let":
		if len(tokens) == 4 && tokens[2] == "=" {
			val, err := strconv.Atoi(tokens[3])
			if err != nil {
				fmt.Println("Invalid value")
				return
			}
			variables[tokens[1]] = val
		}
	case "add":
		if len(tokens) == 4 && tokens[2] == "to" {
			if val, ok := variables[tokens[1]]; ok {
				if addend, err := strconv.Atoi(tokens[3]); err == nil {
					variables[tokens[1]] = val + addend
				}
			}
		}
	default:
		fmt.Println("Unknown command")
	}
}

func main() {
	for {
		var input string
		fmt.Print(">> ")
		fmt.Scanln(&input)

		tokens := tokenize(input)
		interpret(tokens)
	}
}
