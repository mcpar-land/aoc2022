package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day10(input string) error {

	answer := 0

	cycles := 0
	buffer := 1

	currentPixel := 0

	doCycle := func() {

		cycles += 1
		if cycles == 20 || (cycles+20)%40 == 0 {
			signal := cycles * buffer
			answer += signal
		}
		if cycles%40 == 0 {
			fmt.Print("\n")
		} else {
			p := buffer - currentPixel
			if p == 0 || p == 1 || p == -1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		currentPixel += 1
		if currentPixel >= 40 {
			currentPixel = 0
		}
	}

	lines := strings.Split(input, "\n")

	currentLine := 0

	for {
		line := lines[currentLine]
		args := strings.Split(line, " ")
		switch args[0] {
		case "noop":
			doCycle()
			break
		case "addx":
			amt, err := strconv.Atoi(args[1])
			if err != nil {
				panic(err)
			}
			doCycle()
			doCycle()
			buffer += amt
			break
		default:
			panic(fmt.Sprintf("Unrecognized cmd: %s", line))
		}

		currentLine += 1
		if currentLine >= len(lines) {
			currentLine = 0
		}

		if cycles > 621 {
			break
		}
	}
	fmt.Print("\n")

	fmt.Println("Answer 1:", answer)

	return nil
}
