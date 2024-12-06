package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "Day 5: Print Queue/input.txt"
)

func main() {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println(newPuzzle(input).countCorrectlyOrdered())
}

type puzzle struct {
	rules   rules
	updates []update
}

func newPuzzle(input []byte) puzzle {
	return puzzle{
		rules:   newRules(input),
		updates: newUpdates(input),
	}
}

func (p puzzle) countCorrectlyOrdered() int64 {
	count := int64(0)
	for _, u := range p.updates {
		if u.isCorrectlyOrdered(p.rules) {
			count += u[len(u)/2]
		}
	}
	return count
}

type rules struct {
	ruleSet    []rule
	ruleMapByA map[int64][]rule
	ruleMapByB map[int64][]rule
}

func newRules(input []byte) rules {
	var r []rule
	// parse until the newline in the file
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			break
		}
		r = append(r, newRule(line))
	}

	// make a map of rules by a and b
	ruleMapByA := make(map[int64][]rule)
	ruleMapByB := make(map[int64][]rule)
	for _, rule := range r {
		ruleMapByA[rule.a] = append(ruleMapByA[rule.a], rule)
		ruleMapByB[rule.b] = append(ruleMapByB[rule.b], rule)
	}

	return rules{
		ruleSet:    r,
		ruleMapByA: ruleMapByA,
		ruleMapByB: ruleMapByB,
	}
}

type rule struct {
	a int64
	b int64
}

func newRule(input string) rule {
	var a, b int64
	_, err := fmt.Sscanf(input, "%d|%d", &a, &b)
	if err != nil {
		panic(err)
	}
	return rule{a, b}
}

type updates []update

func newUpdates(input []byte) updates {
	var u updates
	// parse from the newline in the file
	parsing := false
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			parsing = true
			continue
		}
		if !parsing {
			continue
		}
		u = append(u, newUpdate(line))
	}
	return u
}

type update []int64

func newUpdate(input string) update {
	// an update is a comma separated list of integers
	var u update
	strs := strings.Split(input, ",")
	for _, str := range strs {
		tmp, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		u = append(u, tmp)
	}
	return u
}

func (u update) isCorrectlyOrdered(r rules) bool {
	for i, a := range u {
		for j, b := range u {
			if i == j {
				continue
			}
			if i < j { // if a comes before b
				for _, rule := range r.ruleMapByB[a] {
					if rule.a == b {
						return false
					}
				}
			} else { // if b comes before a
				for _, rule := range r.ruleMapByA[a] {
					if rule.b == b {
						return false
					}
				}
			}

		}
	}
	return true
}
