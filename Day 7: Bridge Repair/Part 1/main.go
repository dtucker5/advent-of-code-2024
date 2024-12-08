package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("Day 7: Bridge Repair/Part 1/input_example.txt")
	if err != nil {
		panic(err)
	}

	eqs := parseEquations(string(input))
	for _, eq := range eqs {
		fmt.Println(eq)
	}

}

type equations []equation

type equation struct {
	testValue int64
	numbers   []int64
}

func parseEquations(input string) equations {
	var eqs equations
	for _, line := range strings.Split(input, "\n") {
		eqs = append(eqs, parseEquation(line))
	}
	return eqs
}

func parseEquation(input string) equation {
	parts := strings.Split(input, ": ")
	testValue, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		panic(err)
	}
	var numbers []int64
	for _, s := range strings.Split(parts[1], " ") {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return equation{testValue, numbers}
}
