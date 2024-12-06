package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	inputFile = "Day 3: Mull It Over/input.txt"
)

func main() {
	// Read input from file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	// Parse the input
	instructions := newInstructions(input)
	fmt.Println(instructions.exec())
}

type instructions []instruction

func newInstructions(input []byte) instructions {
	matches := regexp.MustCompile(`(mul\(\d+,\d+\))|(don't\(\))|(do\(\))`).FindAll([]byte(input), -1)
	if matches == nil {
		panic("invalid input")
	}
	instructions := make(instructions, len(matches))
	for i, match := range matches {
		instructions[i] = newInstruction(match)
	}
	return instructions
}

func (i instructions) exec() int64 {
	var result int64
	do := true
	for _, instruction := range i {
		if instruction.isDo {
			do = true
			continue
		} else if instruction.isDont {
			do = false
			continue
		} else {
			if do {
				result += instruction.x * instruction.y
			}
		}
	}
	return result
}

type instruction struct {
	isDo   bool
	isDont bool
	x      int64
	y      int64
}

// instructions input will either be "mul(X,Y)" or "don't()" or "do()"
func newInstruction(input []byte) instruction {
	if string(input) == "do()" {
		return instruction{
			isDo: true,
		}
	}
	if string(input) == "don't()" {
		return instruction{
			isDont: true,
		}
	}
	strs := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindSubmatch(input)
	if len(strs) != 3 {
		panic("invalid input")
	}
	x, err := strconv.ParseInt(string(strs[1]), 10, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(string(strs[2]), 10, 64)
	if err != nil {
		panic(err)
	}
	return instruction{
		x: x,
		y: y,
	}
}
