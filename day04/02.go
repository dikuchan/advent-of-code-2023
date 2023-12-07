package day04

import (
	"bufio"
	"strings"
)

func solve2(lines []string) int {
	stash := make(map[int]int, len(lines))
	for i := range lines {
		stash[i+1] = 1
	}
	for _, line := range lines {
		card := parseCard(line)
		wc := card.getWinCount()
		for i := 0; i < stash[card.N]; i++ {
			for j := 0; j < wc; j++ {
				stash[card.N+j+1] += 1
			}
		}
	}
	answer := 0
	for _, n := range stash {
		answer += n
	}
	return answer
}

func (s Solver) Solve2(in string) interface{} {
	reader := strings.NewReader(in)
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return solve2(lines)
}
