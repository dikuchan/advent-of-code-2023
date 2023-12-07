package day01

import (
	"bufio"
	"strings"
	"unicode"

	"github.com/dikuchan/aoc-2023/utils"
)

func solve1Line(line string) int {
	digits := make([]rune, 0)
	for _, c := range line {
		if unicode.IsDigit(c) {
			digits = append(digits, c)
		}
	}
	db := utils.RuneToInteger(digits[0])
	de := utils.RuneToInteger(digits[len(digits)-1])
	return db*10 + de
}

func (s Solver) Solve1(in string) interface{} {
	answer := 0
	reader := strings.NewReader(in)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		answer += solve1Line(line)
	}
	return answer
}
