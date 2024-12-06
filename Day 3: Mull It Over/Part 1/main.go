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
	mulInstructions := newMulInstructions(string(input))
	fmt.Println(mulInstructions.addResults())
}

type mulInstructions []mulInstruction

func newMulInstructions(input string) mulInstructions {
	matches := regexp.MustCompile(`mul\(\d+,\d+\)`).FindAll([]byte(input), -1)
	if matches == nil {
		panic("invalid input")
	}
	instructions := make(mulInstructions, len(matches))
	for i, match := range matches {
		instructions[i] = newMulInstruction(string(match))
	}
	return instructions
}

func (m mulInstructions) addResults() int64 {
	var result int64
	for _, instruction := range m {
		result += instruction.x * instruction.y
	}
	return result
}

type mulInstruction struct {
	x int64
	y int64
}

func newMulInstruction(input string) mulInstruction {
	strs := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindStringSubmatch(input)
	if len(strs) != 3 {
		panic("invalid input")
	}
	x, err := strconv.ParseInt(strs[1], 10, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(strs[2], 10, 64)
	if err != nil {
		panic(err)
	}
	return mulInstruction{
		x: x,
		y: y,
	}
}
