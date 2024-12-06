package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "Day 1: Historian Hysteria/input.txt"
)

func main() {
	// Read input from file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	// Parse the input
	lines := strings.Split(string(input), "\n")
	leftList := make([]int64, len(lines)-1)
	rightList := make([]int64, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		values := strings.Split(line, "   ")
		leftList[i], err = strconv.ParseInt(values[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rightList[i], err = strconv.ParseInt(values[1], 10, 64)
		if err != nil {
			panic(err)
		}
	}

	// Map the right list values by number of occurences
	rightListMap := make(map[int64]int)
	for _, value := range rightList {
		if _, ok := rightListMap[value]; ok {
			rightListMap[value]++
		} else {
			rightListMap[value] = 1
		}
	}

	// Compare the lists
	similarityScore := int64(0)
	for i := 0; i < len(leftList); i++ {
		similarityScore += leftList[i] * int64(rightListMap[leftList[i]])
	}

	fmt.Println(similarityScore)
}
