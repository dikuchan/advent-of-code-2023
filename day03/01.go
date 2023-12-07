package day03

import (
	"strings"
	"unicode"

	"github.com/dikuchan/aoc-2023/utils"
)

func isValid(lines []string, y int, x int) bool {
	N := len(lines)
	M := len(lines[y])

	for _, dy := range []int{-1, 0, 1} {
		j := y + dy
		if j < 0 || j >= N {
			continue
		}
		for _, dx := range []int{-1, 0, 1} {
			i := x + dx
			if dy == 0 && dx == 0 {
				continue
			}
			if i < 0 || i >= M {
				continue
			}
			c := lines[j][i]
			if c != '.' && !utils.IsDigit(c) {
				return true
			}
		}
	}
	return false
}

func (s Solver) Solve1(in string) interface{} {
	var valid, wasNumber bool
	var number = 0
	var numbers = make([]int, 0)

	lines := strings.Split(in, "\n")

	for y, line := range lines {
		for x, c := range line {
			if wasNumber {
				if unicode.IsDigit(c) {
					number = number*10 + utils.RuneToInteger(c)
					if !valid {
						valid = isValid(lines, y, x)
					}
				} else {
					if valid {
						numbers = append(numbers, number)
					}
					wasNumber = false
					valid = false
				}
			} else {
				if unicode.IsDigit(c) {
					number = utils.RuneToInteger(c)
					wasNumber = true
					if !valid {
						valid = isValid(lines, y, x)
					}
				}
			}
		}
	}

	answer := 0
	for _, n := range numbers {
		answer += n
	}
	return answer
}
