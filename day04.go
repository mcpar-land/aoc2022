package main

import (
	"fmt"
	"strconv"
	"strings"
)

type assignmentRange struct {
	start int
	end   int
}

func parseAR(input string) assignmentRange {
	ns := strings.Split(input, "-")
	start, err := strconv.Atoi(ns[0])
	end, err := strconv.Atoi(ns[1])
	if err != nil {
		panic(err)
	}
	return assignmentRange{start, end}
}

func (a assignmentRange) contains(b assignmentRange) bool {
	return b.start >= a.start && b.end <= a.end
}

func (a assignmentRange) overlaps(b assignmentRange) bool {
	return (b.start >= a.start && b.start <= a.end) || (b.end >= a.start && b.end <= a.end)
}

func Day04(input string) error {

	pairs := [][]assignmentRange{}
	for _, line := range strings.Split(input, "\n") {
		pair := strings.Split(line, ",")
		a, b := parseAR(pair[0]), parseAR(pair[1])
		pairs = append(pairs, []assignmentRange{a, b})
	}

	fullContained := 0
	for _, pair := range pairs {
		if pair[0].contains(pair[1]) || pair[1].contains(pair[0]) {
			fullContained += 1
		}
	}

	fmt.Println("Answer 1:", fullContained)

	overlapped := 0
	for _, pair := range pairs {
		if pair[0].overlaps(pair[1]) || pair[1].overlaps(pair[0]) {
			overlapped += 1
		}
	}

	fmt.Println("Answer 2:", overlapped)

	return nil
}
