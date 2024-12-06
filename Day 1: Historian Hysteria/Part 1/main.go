package main

import (
	"fmt"
	"os"
	"sort"
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

	// Sort the lists (using built in sort int func)
	sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })
	sort.Slice(rightList, func(i, j int) bool { return rightList[i] < rightList[j] })

	// Compare the lists
	totalDistance := int64(0)
	for i := 0; i < len(leftList); i++ {
		distance := rightList[i] - leftList[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}

	fmt.Println(totalDistance)
}
