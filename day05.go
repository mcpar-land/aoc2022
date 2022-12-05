package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	moveN    int
	moveFrom int
	moveTo   int
}

var instrFind = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func ParseInstruction(line string) Instruction {
	match := instrFind.FindAllStringSubmatch(line, 4)[0]

	moveN, err := strconv.Atoi(match[1])
	moveFrom, err := strconv.Atoi(match[2])
	moveTo, err := strconv.Atoi(match[3])
	if err != nil {
		panic(err)
	}
	return Instruction{moveN, moveFrom - 1, moveTo - 1}
}

type Crate string

func Day05(input string) error {

	// split the input between the crate diagram and the instruction list.
	_input := strings.Split(input, "\n\n")
	crateDiagram := strings.Split(_input[0], "\n")
	instructions := strings.Split(_input[1], "\n")

	crateCountLine := crateDiagram[len(crateDiagram)-1]
	crateLines := crateDiagram[0 : len(crateDiagram)-1]

	// each crate column is 4 characters wide, except for the last one which is 3
	nCrates := (len(crateCountLine) + 1) / 4

	crates1 := make([][]Crate, nCrates)
	crates2 := make([][]Crate, nCrates)

	// we need to loop through the crate lines backwards to make the stacks
	// correctly.
	for i := len(crateLines) - 1; i >= 0; i-- {
		line := crateLines[i]
		for j := 0; j < nCrates; j++ {
			a := j * 4
			b := a + 4
			if b > len(line) {
				b = len(line)
			}
			cr := strings.Trim(line[a:b], "[] ")
			if cr != "" {
				crates1[j] = append(crates1[j], Crate(cr))
				crates2[j] = append(crates2[j], Crate(cr))
			}
		}
	}

	for _, line := range instructions {
		instr := ParseInstruction(line)
		for i := 0; i < instr.moveN; i++ {
			fromLen := len(crates1[instr.moveFrom])
			crate := crates1[instr.moveFrom][fromLen-1]
			crates1[instr.moveFrom] = crates1[instr.moveFrom][:fromLen-1]
			crates1[instr.moveTo] = append(crates1[instr.moveTo], crate)
		}
	}

	answer := ""
	for _, stack := range crates1 {
		answer += string(stack[len(stack)-1])
	}

	fmt.Println("Answer 1:", answer)

	for _, line := range instructions {
		instr := ParseInstruction(line)
		fromLen := len(crates2[instr.moveFrom])
		crates := crates2[instr.moveFrom][fromLen-instr.moveN:]
		crates2[instr.moveFrom] = crates2[instr.moveFrom][:fromLen-instr.moveN]
		crates2[instr.moveTo] = append(crates2[instr.moveTo], crates...)
	}

	answer = ""
	for _, stack := range crates2 {
		answer += string(stack[len(stack)-1])
	}

	fmt.Println("Answer 2:", answer)

	return nil
}
