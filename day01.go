package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day01(input string) error {

	elves := []int{}
	var elf int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		cals, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		elf += cals
	}
	elves = append(elves, elf)

	sort.Ints(elves)

	fmt.Printf("%v\n", elves)

	fmt.Println("Answer 1:", elves[len(elves)-1])

	top3 := elves[len(elves)-3:]

	fmt.Println(top3)

	totalTop3 := 0

	for _, cals := range top3 {
		totalTop3 += cals
	}

	fmt.Println("Answer 2:", totalTop3)

	return nil
}
