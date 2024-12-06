package main

import (
	"fmt"
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

const (
	maxDampening = 1
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
		safe := report.isSafe(0)
		if safe {
			count++
		}
	}
	return count
}

type report []int64

func newReport(input string) report {
	values := strings.Split(input, " ")
	r := make(report, 0, len(values))
	for _, value := range values {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, int64(i))
	}
	return r
}

func (r report) isSafe(dampeningLevel int) bool {
	if dampeningLevel > maxDampening {
		return false
	}
	//space := ""
	//for i := 0; i < dampeningLevel; i++ {
	//	space += "    "
	//}

	dir := r.getDirection()

	safe := true
	for i := 0; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			safe = false
			break
		}
		if dir == increasing && r[i] > r[i+1] {
			safe = false
			break
		}
		if dir == decreasing && r[i] < r[i+1] {
			safe = false
			break
		}
		if abs(r[i]-r[i+1]) > 3 {
			safe = false
			break
		}
	}

	if safe {
		return true
	}

	for i := 0; i < len(r); i++ {
		r2 := r.without(i)
		safe = r2.isSafe(dampeningLevel + 1)
		if safe {
			return true
		}
	}

	return false
}

func (r report) getDirection() direction {
	if len(r) < 2 || r[0] < r[1] {
		return increasing
	}
	return decreasing
}

func (r report) without(i int) report {
	withoutI := make(report, 0, len(r)-1)
	withoutI = append(withoutI, r[:i]...)
	withoutI = append(withoutI, r[i+1:]...)
	return withoutI
}

//func (r report) without(i int) report {
//	if i < 0 || i >= len(r) {
//		return r
//	}
//	withoutI := make(report, 0, len(r)-1)
//	withoutI = append(withoutI, r[:i]...)
//	if i < len(r)-1 {
//		withoutI = append(withoutI, r[i+1:]...)
//	}
//	return withoutI
//}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
