package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Reverse[T any](input []T) []T {
	if len(input) == 0 {
		return input
	}
	return append(Reverse(input[1:]), input[0])
}

type Grid[T any] [][]T

func NewGrid[T any](width, height int) Grid[T] {
	g := [][]T{}
	for i := 0; i < height; i++ {
		g = append(g, make([]T, width))
	}
	return g
}

func (g Grid[T]) GetRow(i int, reverse bool) []T {
	if reverse {
		return Reverse(g[i])
	}
	return g[i]
}

func (g Grid[T]) GetColumn(i int, reverse bool) []T {
	col := []T{}
	for _, row := range g {
		col = append(col, row[i])
	}
	if reverse {
		return Reverse(col)
	}
	return col
}

func ParseIntGrid(input string) Grid[int] {
	_grid := [][]int{}

	for _, rowStr := range strings.Split(input, "\n") {
		row := []int{}
		for _, cell := range strings.Split(rowStr, "") {
			c, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			row = append(row, c)
		}
		_grid = append(_grid, row)
	}
	return Grid[int](_grid)
}

func Visibility(trees []int) []bool {
	tallest := -1
	visibility := []bool{}
	for _, tree := range trees {
		visibility = append(visibility, tree > tallest)
		if tree > tallest {
			tallest = tree
		}
	}
	return visibility
}

func ScenicScore(g [][]int, x, y int) int {
	a, b, c, d := scenicScoreDir(g, x, y, 1, 0),
		scenicScoreDir(g, x, y, -1, 0),
		scenicScoreDir(g, x, y, 0, 1),
		scenicScoreDir(g, x, y, 0, -1)
	score := a * b * c * d
	// fmt.Printf("Score @ %d,%d is %d * %d * %d * %d = %d\n", x, y, a, b, c, d, score)
	return score
}

func scenicScoreDir(g [][]int, x, y, dx, dy int) int {
	base := g[y][x]
	score := 0
	for {
		x += dx
		y += dy
		if x < 0 || y < 0 || x >= len(g[0]) || y >= len(g) {
			return score
		}
		if g[y][x] >= base {
			score++
			return score
		}
		score++
	}
}

func Day08(input string) error {

	grid := ParseIntGrid(input)

	width, height := len(grid[0]), len(grid)

	visib := NewGrid[bool](width, height)

	for x := 0; x < width; x++ {
		for y, vis := range Visibility(grid.GetColumn(x, false)) {
			if vis {
				visib[y][x] = true
			}
		}
		for y, vis := range Visibility(grid.GetColumn(x, true)) {
			if vis {
				visib[height-y-1][x] = true
			}
		}
	}

	for y := 0; y < height; y++ {
		for x, vis := range Visibility(grid.GetRow(y, false)) {
			if vis {
				visib[y][x] = true
			}
		}
		for x, vis := range Visibility(grid.GetRow(y, true)) {
			if vis {
				visib[y][width-x-1] = true
			}
		}
	}

	count := 0
	for _, row := range visib {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}

	fmt.Println("Answer 1:", count)

	bestScore := 0
	for y, row := range grid {
		for x := range row {
			score := ScenicScore(grid, x, y)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Println("Answer 2:", bestScore)

	return nil
}
