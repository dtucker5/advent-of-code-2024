package main

import (
	"math"
	"testing"
)

func Test_isSafe(t *testing.T) {
	type fields struct {
		report []int64
	}
	tests := []struct {
		name   string
		report report
		want   bool
	}{
		{name: "safely increasing levels", report: []int64{1, 2, 3, 4, 5}, want: true},
		{name: "dangerously increasing levels", report: []int64{1, 2, 3, 7, 8}, want: false},
		{name: "dangerously increasing levels (dampened)", report: []int64{1, 2, 3, 4, 8}, want: true},
		{name: "safely decreasing levels", report: []int64{5, 4, 3, 2, 1}, want: true},
		{name: "dangerously decreasing levels", report: []int64{5, 4, 3, -1, -2}, want: false},
		{name: "dangerously decreasing levels (dampened)", report: []int64{5, 4, 3, 2, -2}, want: true},
		{name: "changing directions", report: []int64{1, 2, 1, 4, 3}, want: false},
		{name: "changing directions (dampened)", report: []int64{1, 2, 1, 4, 5}, want: true},
		{name: "unchanging levels", report: []int64{1, 1, 1, 4, 5}, want: false},
		{name: "unchanging levels (dampened)", report: []int64{1, 1, 3, 4, 5}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.report.isSafe(0); got != tt.want {
				t.Errorf("isSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

type direction int

const (
	increasing direction = iota
	decreasing
)

const (
	maxDampening = 1
)

type report []int64

func (r report) isSafe(dampeningLevel int) bool {
	if dampeningLevel > maxDampening {
		return false
	}

	//fmt.Println(" ", r, dampeningLevel)

	if len(r) < 2 {
		return true
	}

	if r[0] == r[1] {
		return r.without(0).isSafe(dampeningLevel+1) || r.without(1).isSafe(dampeningLevel+1)
	}

	var dir direction
	if r[0] < r[1] {
		dir = increasing
	} else {
		dir = decreasing
	}

	if int64(math.Abs(float64(r[0]-r[1]))) > 3 {
		return r.without(0).isSafe(dampeningLevel+1) || r.without(1).isSafe(dampeningLevel+1)
	}

	for i := 1; i < len(r)-1; i++ {
		if r[i] == r[i+1] {
			return r.without(i).isSafe(dampeningLevel+1) || r.without(i+1).isSafe(dampeningLevel+1)
		}
		if dir == increasing && r[i] > r[i+1] {
			return r.without(i).isSafe(dampeningLevel+1) || r.without(i+1).isSafe(dampeningLevel+1)
		}
		if dir == decreasing && r[i] < r[i+1] {
			return r.without(i).isSafe(dampeningLevel+1) || r.without(i+1).isSafe(dampeningLevel+1)
		}
		if int64(math.Abs(float64(r[i]-r[i+1]))) > 3 {
			return r.without(i).isSafe(dampeningLevel+1) || r.without(i+1).isSafe(dampeningLevel+1)
		}
	}

	return true
}

func (r report) without(i int) report {
	withoutI := make(report, 0, len(r)-1)
	withoutI = append(withoutI, r[:i]...)
	withoutI = append(withoutI, r[i+1:]...)
	return withoutI
}
