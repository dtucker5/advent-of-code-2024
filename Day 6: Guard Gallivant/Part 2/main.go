package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "Day 6: Guard Gallivant/Part 2/input.txt"
)

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(input), "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}

	visited, _ := patrol(grid)

	loopCount := int64(0)
	for _, value := range visited {
		// Copy the grid and put an obstacle at the position identified by the value
		gridCopy := make([][]rune, len(grid))
		for i, row := range grid {
			gridCopy[i] = make([]rune, len(row))
			copy(gridCopy[i], row)
		}
		gridCopy[value.x][value.y] = '#'

		// Check if there is a loop
		_, loop := patrol(gridCopy)
		if loop {
			loopCount++
		}
	}
	fmt.Println("loop count:", loopCount)
}

type positionAndDirection struct {
	x, y      int
	direction rune
}

func patrol(grid [][]rune) (map[string]positionAndDirection, bool) {
	var x, y int
	for i, row := range grid {
		for j, cell := range row {
			if cell == '<' || cell == '^' || cell == '>' || cell == 'v' {
				x, y = i, j
				break
			}
		}
	}

	gridSizeX := len(grid)
	gridSizeY := len(grid[0])

	direction := grid[x][y]

	visited := make(map[string]positionAndDirection)

	// While the guard is in the mapped area
	for {
		if value, found := visited[fmt.Sprintf("%d,%d", x, y)]; found && value.direction == direction {
			// Do not count it as a distinct position
			//fmt.Println("Loop detected at", x, y, direction)
			//fmt.Println("Visited:", visited)
			return visited, true
		}

		visited[fmt.Sprintf("%d,%d", x, y)] = positionAndDirection{x, y, direction}

		for willHitObstacle(grid, x, y, direction) {
			direction = turnRight(direction)
		}

		var ok bool
		if x, y, ok = moveForward(direction, x, y, gridSizeX, gridSizeY); !ok {
			break
		}
	}

	return visited, false
}

func willHitObstacle(grid [][]rune, x, y int, direction rune) bool {
	if direction == '<' {
		if y-1 < 0 {
			return false
		}
		return grid[x][y-1] == '#'
	} else if direction == '^' {
		if x-1 < 0 {
			return false
		}
		return grid[x-1][y] == '#'
	} else if direction == '>' {
		if y+1 >= len(grid[0]) {
			return false
		}
		return grid[x][y+1] == '#'
	} else if direction == 'v' {
		if x+1 >= len(grid) {
			return false
		}
		return grid[x+1][y] == '#'
	}
	return false
}

func turnRight(direction rune) rune {
	switch direction {
	case '<':
		return '^'
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	}
	return direction
}

func moveForward(direction rune, x, y, sizeX, sizeY int) (int, int, bool) {
	switch direction {
	case '<':
		return x, y - 1, y-1 >= 0
	case '^':
		return x - 1, y, x-1 >= 0
	case '>':
		return x, y + 1, y+1 < sizeX
	case 'v':
		return x + 1, y, x+1 < sizeY
	default:
		return x, y, false
	}
}
