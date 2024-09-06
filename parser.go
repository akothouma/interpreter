package main

import (
	"fmt"
	"strings"
)

type Parser struct{}

func (p *Parser) Parse(tokens []string) (Command, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("empty input")
	}

	switch tokens[0] {
	case "var":
		return &VarCommand{tokens}, nil
	case "fmt.Println":
		return &PrintCommand{tokens}, nil
	case "append":
		return &AppendCommand{tokens}, nil
	default:
		if strings.Contains(tokens[0], "[") && len(tokens) >= 3 && tokens[1] == "=" {
			return &SliceAssignCommand{tokens}, nil
		} else if len(tokens) >= 3 && tokens[1] == "=" {
			return &AssignCommand{tokens}, nil
		}
	}

	return nil, fmt.Errorf("unknown command: %s", tokens[0])
}

type Command interface {
	Execute(mm *MemoryManager) error
}