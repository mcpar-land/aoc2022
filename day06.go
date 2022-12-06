package main

import "fmt"

func countOf(input string, v rune) int {
	count := 0
	for _, i := range input {
		if i == v {
			count++
		}
	}
	return count
}

func allDifferent(input string) bool {
	for _, v := range input {
		if countOf(input, v) > 1 {
			return false
		}
	}
	return true
}

func nUniquePos(input string, n int) int {
	for i := 0; i < len(input)-n; i++ {
		m := input[i : i+n]
		if allDifferent(m) {
			return i + n
		}
	}
	return -1
}

func Day06(input string) error {

	fmt.Println("Answer 1:", nUniquePos(input, 4))

	fmt.Println("Answer 2:", nUniquePos(input, 14))

	return nil
}
