package main

import (
	"fmt"
	"strconv"
	"strings"
)

type VarCommand struct {
	tokens []string
}

func (c *VarCommand) Execute(mm *MemoryManager) error {
	if len(c.tokens) < 4 || c.tokens[2] != "=" {
		return fmt.Errorf("invalid var syntax")
	}

	name := c.tokens[1]

	if c.tokens[3] == "[]int" {
		mm.SetSlice(name, []int{})
	} else if strings.HasPrefix(c.tokens[3], "[]int{") && strings.HasSuffix(c.tokens[len(c.tokens)-1], "}") {
		values := []int{}
		for _, v := range c.tokens[4 : len(c.tokens)-1] {
			v = strings.Trim(v, ",}")
			val, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("invalid value in slice declaration: %s", v)
			}
			values = append(values, val)
		}
		mm.SetSlice(name, values)
	} else {
		val, err := strconv.Atoi(c.tokens[3])
		if err != nil {
			return fmt.Errorf("invalid value: %s", c.tokens[3])
		}
		mm.SetVariable(name, val)
	}

	return nil
}

type PrintCommand struct {
	tokens []string
}

func (c *PrintCommand) Execute(mm *MemoryManager) error {
	if len(c.tokens) < 2 {
		return fmt.Errorf("print command requires an argument")
	}

	arg := c.tokens[1]
	if strings.Contains(arg, "[") {
		parts := strings.SplitN(arg, "[", 2)
		sliceName := parts[0]
		index, err := strconv.Atoi(strings.Trim(parts[1], "[]"))
		if err != nil {
			return fmt.Errorf("invalid index: %s", parts[1])
		}

		slice, ok := mm.GetSlice(sliceName)
		if !ok {
			return fmt.Errorf("slice not found: %s", sliceName)
		}

		if index < 0 || index >= len(slice) {
			return fmt.Errorf("index out of range for slice: %s", sliceName)
		}

		fmt.Println(slice[index])
	} else if val, ok := mm.GetVariable(arg); ok {
		fmt.Println(val)
	} else if slice, ok := mm.GetSlice(arg); ok {
		fmt.Println(slice)
	} else {
		fmt.Println(arg)
	}

	return nil
}

type AppendCommand struct {
	tokens []string
}

func (c *AppendCommand) Execute(mm *MemoryManager) error {
	if len(c.tokens) != 4 || c.tokens[2] != "=" {
		return fmt.Errorf("invalid append syntax")
	}

	sliceName := c.tokens[1]
	slice, ok := mm.GetSlice(sliceName)
	if !ok {
		return fmt.Errorf("slice not found: %s", sliceName)
	}

	val, err := strconv.Atoi(c.tokens[3])
	if err != nil {
		return fmt.Errorf("invalid value to append: %s", c.tokens[3])
	}

	mm.SetSlice(sliceName, append(slice, val))
	return nil
}

type SliceAssignCommand struct {
	tokens []string
}

func (c *SliceAssignCommand) Execute(mm *MemoryManager) error {
	parts := strings.SplitN(c.tokens[0], "[", 2)
	sliceName := parts[0]
	index, err := strconv.Atoi(strings.Trim(parts[1], "[]"))
	if err != nil {
		return fmt.Errorf("invalid index: %s", parts[1])
	}

	slice, ok := mm.GetSlice(sliceName)
	if !ok {
		return fmt.Errorf("slice not found: %s", sliceName)
	}

	if index < 0 || index >= len(slice) {
		return fmt.Errorf("index out of range for slice: %s", sliceName)
	}

	val, err := strconv.Atoi(c.tokens[2])
	if err != nil {
		return fmt.Errorf("invalid value: %s", c.tokens[2])
	}

	slice[index] = val
	mm.SetSlice(sliceName, slice)
	return nil
}

type AssignCommand struct {
	tokens []string
}

func (c *AssignCommand) Execute(mm *MemoryManager) error {
	varName := c.tokens[0]
	_, ok := mm.GetVariable(varName)
	if !ok {
		return fmt.Errorf("variable not declared: %s", varName)
	}

	if len(c.tokens) == 3 {
		val, err := strconv.Atoi(c.tokens[2])
		if err != nil {
			return fmt.Errorf("invalid value: %s", c.tokens[2])
		}
		mm.SetVariable(varName, val)
	} else if len(c.tokens) == 5 && c.tokens[3] == "+" {
		val1, _ := mm.GetVariable(varName)
		val2, err := strconv.Atoi(c.tokens[4])
		if err != nil {
			return fmt.Errorf("invalid value: %s", c.tokens[4])
		}
		mm.SetVariable(varName, val1+val2)
	} else {
		return fmt.Errorf("invalid assignment syntax")
	}

	return nil
}