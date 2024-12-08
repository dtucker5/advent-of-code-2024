package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("Day 7: Bridge Repair/Part 2/input.txt")
	if err != nil {
		panic(err)
	}

	eqs := parseEquations(string(input))

	count := int64(0)
	totalCalibrationResult := int64(0)
	ops := []rune{'+', '*', '|'}
	for _, eq := range eqs {
		if eq.canBeSolvedWithOperators(ops) {
			count++
			totalCalibrationResult += eq.testValue
		}
	}
	fmt.Println(count)
	fmt.Println(totalCalibrationResult)
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

func (eq equation) canBeSolvedWithOperators(ops []rune) bool {
	solvable := false
	solveEquation(eq.numbers, ops, func(result int64) {
		if result == eq.testValue {
			solvable = true
		}
	})
	return solvable
}

func solveEquation(numbers []int64, ops []rune, callback func(int64)) {
	var generateCombinations func([]int64, []rune)
	generateCombinations = func(nums []int64, currentOps []rune) {
		if len(currentOps) == len(nums)-1 {
			result := nums[0]
			for i, op := range currentOps {
				switch op {
				case '+':
					result += nums[i+1]
				case '*':
					result *= nums[i+1]
				case '|':
					var err error
					result, err = strconv.ParseInt(strconv.FormatInt(result, 10)+strconv.FormatInt(nums[i+1], 10), 10, 64)
					if err != nil {
						panic(err)
					}
				}
			}
			callback(result)
			return
		}
		for _, op := range ops {
			generateCombinations(nums, append(currentOps, op))
		}
	}
	generateCombinations(numbers, []rune{})
}
