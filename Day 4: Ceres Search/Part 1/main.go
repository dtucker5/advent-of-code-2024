package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	inputFile = "Day 4: Ceres Search/input.txt"
)

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(newWordSearch(input).countWordOccurrences("XMAS"))
}

type wordSearch struct {
	grid       [][]byte
	heuristics []func([][]byte, string) int64
}

func newWordSearch(input []byte) *wordSearch {
	grid := make([][]byte, 0)
	rows := strings.Split(string(input), "\n")
	for _, line := range rows {
		grid = append(grid, []byte(line))
	}
	return &wordSearch{
		grid: grid,
		heuristics: []func([][]byte, string) int64{
			countWordOccurrencesByRow,
			countWordOccurrencesByColumn,
			countWordOccurrencesByDiagonal,
		},
	}
}

func (ws *wordSearch) countWordOccurrences(word string) int64 {
	count := int64(0)
	for _, heuristic := range ws.heuristics {
		count += heuristic(ws.grid, word)
	}
	return count
}

func countWordOccurrencesByRow(grid [][]byte, word string) int64 {
	count := int64(0)
	r1 := regexp.MustCompile(word)
	r2 := regexp.MustCompile(reverseString(word))
	for _, row := range grid {
		// left to right
		found := r1.FindAll(row, -1)
		count += int64(len(found))
		// right to left
		found = r2.FindAll(row, -1)
		count += int64(len(found))
	}
	return count
}

func countWordOccurrencesByColumn(grid [][]byte, word string) int64 {
	count := int64(0)
	r1 := regexp.MustCompile(word)
	r2 := regexp.MustCompile(reverseString(word))
	for i := 0; i < len(grid); i++ {
		column := make([]byte, len(grid))
		for j := 0; j < len(grid); j++ {
			column[j] = grid[j][i]
		}
		// top to bottom
		found := r1.FindAll(column, -1)
		count += int64(len(found))
		// bottom to top
		found = r2.FindAll(column, -1)
		count += int64(len(found))
	}
	return count
}

func countWordOccurrencesByDiagonal(grid [][]byte, word string) int64 {
	count := int64(0)
	r1 := regexp.MustCompile(word)
	r2 := regexp.MustCompile(reverseString(word))

	// To do this we will iterate from x = 0 to x = len(grid) - len(word) and y = 0 to y = len(grid) - len(word) to cover the entire grid
	// At each iteration we will check the diagonal from the current x and y position to the bottom right corner

	sizeX := len(grid)
	sizeY := len(grid[0])

	// downwards from the top and left

	// across the top
	for i := 0; i < len(grid); i++ {
		// across the top
		s := ""
		for x, y := i, 0; x < sizeX && y < sizeY; x, y = x+1, y+1 {
			s += string(grid[y][x])
		}
		// check if the diagonal contains the word
		found := r1.FindAll([]byte(s), -1)
		count += int64(len(found))
		// check the reverse diagonal
		found = r2.FindAll([]byte(s), -1)
		count += int64(len(found))
	}

	// down the left
	for i := 1; i < len(grid[0]); i++ {
		// down the left
		s := ""
		for x, y := 0, i; x < sizeX && y < sizeY; x, y = x+1, y+1 {
			s += string(grid[y][x])
		}
		// check if the diagonal contains the word
		found := r1.FindAll([]byte(s), -1)
		count += int64(len(found))
		// check the reverse diagonal
		found = r2.FindAll([]byte(s), -1)
		count += int64(len(found))
	}

	// upwards from the bottom and left

	// across the bottom
	for i := 0; i < len(grid); i++ {
		// across the bottom
		s := ""
		for x, y := i, sizeY-1; x < sizeX && y >= 0; x, y = x+1, y-1 {
			s += string(grid[y][x])
		}
		// check if the diagonal contains the word
		found := r1.FindAll([]byte(s), -1)
		count += int64(len(found))
		// check the reverse diagonal
		found = r2.FindAll([]byte(s), -1)
		count += int64(len(found))
	}

	// up the left
	for i := sizeY - 2; i >= 0; i-- {
		// up the left
		s := ""
		for x, y := 0, i; x < sizeX && y >= 0; x, y = x+1, y-1 {
			s += string(grid[y][x])
		}
		// check if the diagonal contains the word
		found := r1.FindAll([]byte(s), -1)
		count += int64(len(found))
		// check the reverse diagonal
		found = r2.FindAll([]byte(s), -1)
		count += int64(len(found))
	}

	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
