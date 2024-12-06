package main

import (
	"fmt"
	"os"
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

	fmt.Println(newWordSearch(input).countWordOccurrences())
}

type wordSearch struct {
	grid       [][]byte
	heuristics []func([][]byte) int64
}

func newWordSearch(input []byte) *wordSearch {
	grid := make([][]byte, 0)
	rows := strings.Split(string(input), "\n")
	for _, line := range rows {
		grid = append(grid, []byte(line))
	}
	return &wordSearch{
		grid: grid,
		heuristics: []func([][]byte) int64{
			countXMasOccurrences,
		},
	}
}

func (ws *wordSearch) countWordOccurrences() int64 {
	count := int64(0)
	for _, heuristic := range ws.heuristics {
		count += heuristic(ws.grid)
	}
	return count
}

func countXMasOccurrences(grid [][]byte) int64 {
	count := int64(0)
	sizeX := len(grid[0])
	sizeY := len(grid)
	for y := 0; y < sizeY-2; y++ {
		for x := 0; x < sizeX-2; x++ {
			d1 := make([]byte, 3)
			d1[0] = grid[y][x]
			d1[1] = grid[y+1][x+1]
			d1[2] = grid[y+2][x+2]
			d2 := make([]byte, 3)
			d2[0] = grid[y][x+2]
			d2[1] = grid[y+1][x+1]
			d2[2] = grid[y+2][x]
			if (string(d1) == "MAS" || string(d1) == "SAM") && (string(d2) == "MAS" || string(d2) == "SAM") {
				count++
			}
		}
	}
	return count
}
