package day03

import (
	"strings"
	"unicode"

	"github.com/dikuchan/aoc-2023/utils"
)

type P struct {
	X int
	Y int
}

func NewP(y int, x int) P {
	return P{y, x}
}

func findAsterisks(lines []string, y int, x int) []P {
	N := len(lines)
	M := len(lines[y])

	asterisks := make([]P, 0)

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
			if c == '*' {
				asterisks = append(asterisks, NewP(j, i))
			}
		}
	}
	return asterisks
}

func (s Solver) Solve2(in string) interface{} {
	var wasNumber bool
	var number = 0

	var numbers = make(map[int]map[int][]int)
	var asterisks = make(map[P]struct{}, 0)

	lines := strings.Split(in, "\n")

	for y, line := range lines {
		for x, c := range line {
			if wasNumber {
				if unicode.IsDigit(c) {
					number = number*10 + utils.RuneToInteger(c)
					for _, p := range findAsterisks(lines, y, x) {
						asterisks[p] = struct{}{}
					}
				} else {
					for p := range asterisks {
						if _, ok := numbers[p.Y]; !ok {
							numbers[p.Y] = make(map[int][]int)
						}
						if _, ok := numbers[p.Y][p.X]; !ok {
							numbers[p.Y][p.X] = make([]int, 0)
						}
						numbers[p.Y][p.X] = append(numbers[p.Y][p.X], number)
					}
					wasNumber = false
					asterisks = make(map[P]struct{}, 0)
				}
			} else {
				if unicode.IsDigit(c) {
					number = utils.RuneToInteger(c)
					wasNumber = true
					for _, p := range findAsterisks(lines, y, x) {
						asterisks[p] = struct{}{}
					}
				}
			}
		}
	}

	answer := 0
	for _, ys := range numbers {
		for _, xs := range ys {
			if len(xs) == 2 {
				answer += xs[0] * xs[1]
			}
		}
	}
	return answer
}
