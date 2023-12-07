package day01

import (
	"bufio"
	"strings"
	"unicode"

	"github.com/dikuchan/aoc-2023/utils"
)

var NUMBERS = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func solve2Line(line string) int {
	digits := make([]rune, 0)
	for i, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		} else {
			for s, n := range NUMBERS {
				if strings.HasPrefix(line[i:], s) {
					digits = append(digits, n)
				}
			}
		}
	}
	db := utils.RuneToInteger(digits[0])
	de := utils.RuneToInteger(digits[len(digits)-1])
	return db*10 + de
}

func (s Solver) Solve2(in string) interface{} {
	answer := 0
	reader := strings.NewReader(in)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		answer += solve2Line(line)
	}
	return answer
}
