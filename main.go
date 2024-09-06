package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	interpreter := NewInterpreter(reader)
	interpreter.Run()
}

type Interpreter struct {
	lexer   *Lexer
	parser  *Parser
	memory  *MemoryManager
	scanner *bufio.Scanner
}

func NewInterpreter(reader *bufio.Reader) *Interpreter {
	return &Interpreter{
		lexer:   &Lexer{},
		parser:  &Parser{},
		memory:  NewMemoryManager(),
		scanner: bufio.NewScanner(reader),
	}
}

func (i *Interpreter) Run() {
	for i.scanner.Scan() {
		line := i.scanner.Text()
		tokens := i.lexer.Tokenize(line)
		command, err := i.parser.Parse(tokens)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		err = command.Execute(i.memory)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}

	if err := i.scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}