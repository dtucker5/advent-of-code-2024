package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "Day 2: Red-Nosed Reports/input.txt"
)

func main() {
	// Read input from file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	// Parse the input
	reports := newReports(input)
	fmt.Println(reports.countSafe())
}

type direction int

const (
	increasing direction = iota
	decreasing
)

type reports []report

func newReports(input []byte) reports {
	lines := strings.Split(string(input), "\n")
	r := make(reports, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		r = append(r, newReport(line))
	}
	return r
}

func (r reports) countSafe() int64 {
	var count int64
	for _, report := range r {
		if report.isSafe() {
			count++
		}
	}
	return count
}

type report []level

func newReport(input string) report {
	values := strings.Split(input, " ")
	r := make(report, 0, len(values))
	for _, value := range values {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, level(i))
	}
	return r
}

func (r report) isSafe() bool {
	if len(r) < 2 {
		return true
	}

	if r[0] == r[1] {
		return false
	}

	var dir direction
	if r[0] < r[1] {
		dir = increasing
	} else {
		dir = decreasing
	}

	if int64(math.Abs(float64(r[0]-r[1]))) > 3 {
		return false
	}

	for i := 1; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			return false
		}
		if dir == increasing && r[i] > r[i+1] {
			return false
		}
		if dir == decreasing && r[i] < r[i+1] {
			return false
		}
		if int64(math.Abs(float64(r[i]-r[i+1]))) > 3 {
			return false
		}
	}

	return true
}

type level int64
