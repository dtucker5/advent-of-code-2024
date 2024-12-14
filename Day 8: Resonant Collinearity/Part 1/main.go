package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("Day 8: Resonant Collinearity/Part 1/input.txt")
	if err != nil {
		panic(err)
	}

	// Calculate the impact of the signal. How many unique locations within the bounds of the map contain an antinode?

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

	// For each pair of antennas with the same frequency, calculate the position of both antinodes.
	// Store all antinodes in a set (to avoid duplicates).
	antinodeSet := make(map[Coordinate]struct{})

	for _, antennas := range antennaLists {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				if i == j {
					continue // Skip if the antennas are the same.
				}
				// Calculate the position of the two antinodes.
				// An antinode occurs at any point that is perfectly in line with two antennas of the same frequency -
				// but only when one of the antennas is twice as far away as the other. This means that for any pair of
				// antennas with the same frequency, there are two antinodes, one on either side of them.

				diffX := antennas[j].X - antennas[i].X
				diffY := antennas[j].Y - antennas[i].Y
				// If the antennas are not in a straight line, skip.
				//if diffX == 0 || diffY == 0 || diffX != diffY {
				//	continue
				//}

				//v1 := Coordinate{
				//	X: antennas[j].X - antennas[i].X,
				//	Y: antennas[j].Y - antennas[i].Y,
				//}
				//v2 := Coordinate{
				//	X: antennas[i].X - antennas[j].X,
				//	Y: antennas[i].Y - antennas[j].Y,
				//}

				// Given two antennas at (x1, y1) and (x2, y2), the antinodes will be at (x1 + (x2 - x1), y1 + (y2 - y1)) and (x2 + (x1 - x2), y2 + (y1 - y2)).
				antinode1 := Coordinate{antennas[j].X + diffX, antennas[j].Y + diffY}
				antinode2 := Coordinate{antennas[i].X - diffX, antennas[i].Y - diffY}

				//fmt.Println("nodes at", antennas[i], antennas[j], "will produce antinodes at", antinode1, antinode2)

				// Add the antinode positions to the set if they are within bounds.
				//antinode1 := Coordinate{antennas[i].X + v1.X, antennas[i].Y + v1.Y}
				//antinode2 := Coordinate{antennas[j].X + v2.X, antennas[j].Y + v2.Y}

				if isWithinBounds(antinode1, len(grid), len(grid[0])) {
					antinodeSet[antinode1] = struct{}{}
					//grid[antinode1.Y][antinode1.X] = '#' // Update the grid with a '#' to indicate the antinode.
				}
				if isWithinBounds(antinode2, len(grid), len(grid[0])) {
					antinodeSet[antinode2] = struct{}{}
					//grid[antinode2.Y][antinode2.X] = '#' // Update the grid with a '#' to indicate the antinode.
				}
			}
		}
	}

	// Ouput the number of unique antinode locations.
	fmt.Println(len(antinodeSet)) // 386 too high

	// Pretty print the grid.
	//for _, row := range grid {
	//	fmt.Println(string(row))
	//}
}

type Coordinate struct {
	X, Y int
}

// Check if the coordinate is within the bounds of the grid.
func isWithinBounds(coord Coordinate, rows, cols int) bool {
	return coord.X >= 0 && coord.X < cols && coord.Y >= 0 && coord.Y < rows
}
