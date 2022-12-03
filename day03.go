package main

import (
	"fmt"
	"strings"
)

const prios = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func priorityOf(item rune) int {
	i := strings.IndexRune(prios, item)
	if i == -1 {
		panic("Invalid item")
	}
	return i + 1
}

type group []rucksack

type rucksack struct {
	a compartment
	b compartment
}

func newRucksack(input string) rucksack {
	l := len(input) / 2
	return rucksack{
		a: newCompartment(input[:l]),
		b: newCompartment(input[l:]),
	}
}

func (r rucksack) commonPrio() int {
	for _, pa := range r.a {
		for _, pb := range r.b {
			if pa == pb {
				return pa
			}
		}
	}
	panic("No common value in rucksack")
}

type compartment map[rune]int

func newCompartment(input string) compartment {
	c := compartment(map[rune]int{})
	for _, v := range input {
		c[v] = priorityOf(v)
	}
	return c
}

func Day03(input string) error {
	sacks := []rucksack{}
	for _, line := range strings.Split(input, "\n") {
		sacks = append(sacks, newRucksack(line))
	}

	sum := 0
	for _, sack := range sacks {
		sum += sack.commonPrio()
	}
	fmt.Println("Answer 1:", sum)

	elves := []compartment{}
	for _, line := range strings.Split(input, "\n") {
		elves = append(elves, newCompartment(line))
	}

	sum = 0
Elves:
	for i := 0; i < len(elves); i += 3 {
		a, b, c := elves[i], elves[i+1], elves[i+2]
		for _, va := range a {
			for _, vb := range b {
				if va == vb {
					for _, vc := range c {
						if vb == vc {
							sum += vc
							continue Elves
						}
					}
				}
			}
		}
	}

	fmt.Println("Answer 2:", sum)

	return nil
}
