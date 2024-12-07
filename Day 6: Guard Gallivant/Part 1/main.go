package main

import (
	"os"
	"strings"
)

const (
	inputFile = "Day 6: Guard Gallivant/Part 1/input.txt"
)

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	// Lab guards in 1518 follow a very strict patrol protocol which involves repeatedly following these steps:
	//   If there is something directly in front of you, turn right 90 degrees.
	//   Otherwise, take a step forward.
	// How many distinct positions will the guard visit before leaving the mapped area?
	// Position of the guard is marked with <, ^, > or v, indicating the direction the guard is facing.
	// Obstacles are marked with #.
	// Open spaces are marked with a period (.).

	// Find starting position (it counts as a distinct position)
	// While the guard is in the mapped area:
	//   If there is something directly in front of the guard, turn right 90 degrees.
	//   Otherwise, take a step forward.
	//   Mark the position of the guard with a distinct character.
	//   If the guard has visited this position before, do not count it as a distinct position.
	// Return the number of distinct positions the guard visited.

	// Convert the input into a 2D array of runes
	rows := strings.Split(string(input), "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}

	// Find the starting position of the guard (can be anywhere on the grid)
	var x, y int
	for i, row := range grid {
		for j, cell := range row {
			if cell == '<' || cell == '^' || cell == '>' || cell == 'v' {
				x, y = i, j
				break
			}
		}
	}

	//println("guard starting position:", x, y)

	// Find the size of the grid
	gridSizeX := len(grid)
	gridSizeY := len(grid[0])

	// Initialize the number of distinct positions the guard visited
	distinctPositions := 1

	// Initialize the direction the guard is facing
	direction := grid[x][y]

	// Mark the starting position of the guard
	grid[x][y] = 'X'

	// While the guard is in the mapped area
	for {
		//println("guard position:", x, y, direction == '<', direction == '^', direction == '>', direction == 'v')

		// If there is something directly in front of the guard
		if direction == '<' {
			// Check if there is something directly in front of the guard
			if y-1 < gridSizeY && grid[x][y-1] == '#' {
				//println("guard is facing left and there is something directly in front of it")
				direction = '^' // Turn right 90 degrees
				continue
			}
		} else if direction == '^' {
			// Check if there is something directly in front of the guard
			if x-1 < gridSizeX && grid[x-1][y] == '#' {
				//println("guard is facing up and there is something directly in front of it")
				direction = '>' // Turn right 90 degrees
				continue
			}
		} else if direction == '>' {
			// Check if there is something directly in front of the guard
			if y+1 < gridSizeY && grid[x][y+1] == '#' {
				//println("guard is facing right and there is something directly in front of it")
				direction = 'v' // Turn right 90 degrees
				continue
			}
		} else if direction == 'v' {
			// Check if there is something directly in front of the guard
			if x+1 < gridSizeX && grid[x+1][y] == '#' {
				//println("guard is facing down and there is something directly in front of it")
				direction = '<' // Turn right 90 degrees
				continue
			}
		}

		// Check if the guard is at the edge of the grid and take a step forward if not
		if direction == '<' {
			//println("guard is facing left")
			if y-1 < 0 {
				//println("guard is at the edge of the grid")
				break
			}
			y--
		} else if direction == '^' {
			//println("guard is facing up")
			if x-1 < 0 {
				//println("guard is at the edge of the grid (2)")
				break
			}
			x--
		} else if direction == '>' {
			//println("guard is facing right")
			if y+1 >= gridSizeY {
				//println("guard is at the edge of the grid (3)")
				break
			}
			y++
		} else if direction == 'v' {
			//println("guard is facing down")
			if x+1 >= gridSizeX {
				//println("guard is at the edge of the grid (4)")
				break
			}
			x++
		}

		// If the guard has visited this position before
		if grid[x][y] == 'X' {
			// Do not count it as a distinct position
			continue
		}

		// Mark the position of the guard with a distinct character
		grid[x][y] = 'X'

		// Count it as a distinct position
		distinctPositions++
	}

	// Return the number of distinct positions the guard visited
	println(distinctPositions)

}
