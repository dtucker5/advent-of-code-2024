package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("Day 8: Resonant Collinearity/Part 2/input.txt")
	if err != nil {
		panic(err)
	}

	// Convert input into grid.
	rows := strings.Split(string(input), "\n")
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}

	// Store all antennas with the same frequency in individual lists of coordinates.
	antennaLists := make(map[rune][]Coordinate)
	for y, row := range grid {
		for x, frequency := range row {
			if frequency == '.' { // not an antenna
				continue
			}
			antennaLists[frequency] = append(antennaLists[frequency], Coordinate{x, y})
		}
	}

	// Create a set to store unique antinode locations.
	antinodeSet := make(map[Coordinate]struct{})

	for _, antennas := range antennaLists {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Calculate the vector between the two antennas.
				v := Coordinate{
					X: antennas[j].X - antennas[i].X,
					Y: antennas[j].Y - antennas[i].Y,
				}

				// Add the antinode positions to the set if they are within bounds.
				antinode1 := Coordinate{antennas[i].X + v.X, antennas[i].Y + v.Y}
				antinode2 := Coordinate{antennas[j].X - v.X, antennas[j].Y - v.Y}

				if isWithinBounds(antinode1, len(grid), len(grid[0])) {
					antinodeSet[antinode1] = struct{}{}
				}
				if isWithinBounds(antinode2, len(grid), len(grid[0])) {
					antinodeSet[antinode2] = struct{}{}
				}

				// Extrapolate the antinodes until they are out of bounds.
				antinode1.X += v.X
				antinode1.Y += v.Y
				for isWithinBounds(antinode1, len(grid), len(grid[0])) {
					antinodeSet[antinode1] = struct{}{}
					antinode1.X += v.X
					antinode1.Y += v.Y
				}
				antinode2.X -= v.X
				antinode2.Y -= v.Y
				for isWithinBounds(antinode2, len(grid), len(grid[0])) {
					antinodeSet[antinode2] = struct{}{}
					antinode2.X -= v.X
					antinode2.Y -= v.Y
				}
			}
		}
	}

	// Convert the set to a list of unique antinode locations.
	uniqueAntinodes := make([]Coordinate, 0, len(antinodeSet))
	for coord := range antinodeSet {
		uniqueAntinodes = append(uniqueAntinodes, coord)
	}

	// Print the number of unique antinode locations.
	fmt.Println(len(uniqueAntinodes))
}

type Coordinate struct {
	X, Y int
}

// Check if the coordinate is within the bounds of the grid.
func isWithinBounds(coord Coordinate, rows, cols int) bool {
	return coord.X >= 0 && coord.X < cols && coord.Y >= 0 && coord.Y < rows
}
