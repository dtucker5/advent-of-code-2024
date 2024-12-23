package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("Day 10: Hoof It/Part 1/input_example.txt")
	if err != nil {
		panic(err)
	}

	// Parse input into a 2D array
	lines := strings.Split(string(input), "\n")
	numLines := len(lines)
	m := make([][]int, numLines)
	for y, line := range lines {
		m[y] = make([]int, len(line))
		for x, char := range line {
			if char == '.' {
				m[y][x] = -1
			} else {
				i, _ := strconv.Atoi(string(char))
				m[y][x] = i
			}
		}
	}

	prettyPrint(m)

	// Find each trail
	numTrailHeads := 0
	sumTrailScores := 0
	lenY := len(m)
	lenX := len(m[0])
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {
			if m[y][x] == 0 {
				numTrailHeads++
				trailScore := getTrailScore(copyMap(m), x, y)
				println("Trail at (", x, ",", y, ") has a score of", trailScore)
				sumTrailScores += trailScore
			}
		}
	}
	println(numTrailHeads)
	println(sumTrailScores) // 1162 too high
}

func getTrailScore(m [][]int, x, y int) int {
	score := 0

	if m[y][x] == 9 {
		prettyPrint(m)
		return score + 1
	}

	// Check if we can move downwards
	if y+1 < len(m) {
		currentHeight := m[y][x]
		nextHeight := m[y+1][x]
		println(currentHeight, nextHeight)
		if m[y+1][x] == m[y][x]+1 {
			score += getTrailScore(m, x, y+1)
		}
	}

	// Check if we can move upwards
	if y-1 >= 0 && m[y-1][x] == m[y][x]+1 {
		score += getTrailScore(m, x, y-1)
	}

	// Check if we can move rightwards
	if x+1 < len(m[y]) && m[y][x+1] == m[y][x]+1 {
		score += getTrailScore(m, x+1, y)
	}

	// Check if we can move leftwards
	if x-1 >= 0 && m[y][x-1] == m[y][x]+1 {
		score += getTrailScore(m, x-1, y)
	}

	return score
}

func prettyPrint(m [][]int) {
	for _, row := range m {
		for _, cell := range row {
			if cell == -1 {
				fmt.Print(". ")
				continue
			}
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}

func copyMap(m [][]int) [][]int {
	mCopy := make([][]int, len(m))
	for y, row := range m {
		mCopy[y] = make([]int, len(row))
		copy(mCopy[y], row)
	}
	return mCopy
}
