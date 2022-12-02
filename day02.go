package main

import (
	"fmt"
	"strings"
)

type Shape int

const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

func (s Shape) LosesVs() Shape {
	if s == Scissors {
		return Rock
	}
	return s + 1
}

func (s Shape) WinsVs() Shape {
	if s == Rock {
		return Scissors
	}
	return s - 1
}

func (a Shape) ScoreVs(b Shape) int {
	if a == b {
		return 3
	}
	if a.WinsVs() == b {
		return 6
	}
	return 0
}

var shapeLabels = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

type Play int

const (
	Win  Play = 1
	Lose Play = 2
	Draw Play = 3
)

var playLabels = map[string]Play{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

func MoveVs(a Shape, play Play) Shape {
	if play == Draw {
		return a
	} else if play == Lose {
		return a.WinsVs()
	}
	return a.LosesVs()
}

func Day02(input string) error {
	var score int
	var score2 int
	for _, line := range strings.Split(input, "\n") {
		// Part 1
		moves := strings.Split(line, " ")
		a, b := shapeLabels[moves[0]], shapeLabels[moves[1]]
		score += b.ScoreVs(a) + int(b)

		// Part 2
		play := playLabels[moves[1]]
		b = MoveVs(a, play)
		score2 += b.ScoreVs(a) + int(b)
	}
	fmt.Println("Answer 1:", score)
	fmt.Println("Answer 2:", score2)
	return nil
}
