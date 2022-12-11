package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Sign(v int) int {
	if v > 0 {
		return 1
	}
	if v < 0 {
		return -1
	}
	return 0
}

type Pos struct {
	X int
	Y int
}

func (a Pos) Add(b Pos) Pos {
	return Pos{a.X + b.X, a.Y + b.Y}
}

func (a Pos) OffsetSign(b Pos) Pos {
	return Pos{
		X: Sign(a.X - b.X),
		Y: Sign(a.Y - b.Y),
	}
}

func (a Pos) OffsetAbs(b Pos) Pos {
	p := Pos{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
	if p.X < 0 {
		p.X *= -1
	}
	if p.Y < 0 {
		p.Y *= -1
	}
	return p
}

func (tail Pos) MoveToward(head Pos) Pos {
	oa := head.OffsetAbs(tail)
	if oa.X <= 1 && oa.Y <= 1 {
		return tail
	}
	return tail.Add(head.OffsetSign(tail))
}

type Rope []Pos

func NewRope(length int) Rope {
	r := Rope{}
	for i := 0; i < length; i++ {
		r = append(r, Pos{0, 0})
	}
	return r
}

func (r *Rope) Recalculate() {
	l := len(*r)
	for i := 1; i < l; i++ {
		rope := *r
		(*r)[i] = rope[i].MoveToward(rope[i-1])
	}
}

type RopeMove struct {
	Delta Pos
	Amt   int
}

func ParseRopeMove(input string) RopeMove {
	_input := strings.Split(input, " ")
	dir, amt := _input[0], _input[1]
	amtInt, err := strconv.Atoi(amt)
	if err != nil {
		panic(err)
	}
	var delta Pos
	switch dir {
	case "R":
		delta = Pos{1, 0}
		break
	case "L":
		delta = Pos{-1, 0}
		break
	case "U":
		delta = Pos{0, 1}
		break
	case "D":
		delta = Pos{0, -1}
		break
	default:
		panic("Unrecognized command: " + dir)
	}
	return RopeMove{delta, amtInt}
}

func (r Rope) Execute(moves []RopeMove) int {
	visits := map[Pos]bool{}

	visits[Pos{0, 0}] = true

	for _, move := range moves {
		for i := 0; i < move.Amt; i++ {
			r[0] = r[0].Add(move.Delta)
			r.Recalculate()
			visits[r[len(r)-1]] = true
		}
	}

	return len(visits)
}

func Day09(input string) error {

	moves := []RopeMove{}
	for _, line := range strings.Split(input, "\n") {
		moves = append(moves, ParseRopeMove(line))
	}

	fmt.Println("Answer 1:", NewRope(2).Execute(moves))

	fmt.Println("Answer 2:", NewRope(10).Execute(moves))

	return nil
}
