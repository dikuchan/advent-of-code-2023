package day04

import (
	"bufio"
	"strconv"
	"strings"
)

type Card struct {
	N       int
	Winning map[int]struct{}
	My      map[int]struct{}
}

func (c Card) getWinCount() int {
	wc := 0
	for n := range c.My {
		if _, ok := c.Winning[n]; ok {
			wc++
		}
	}
	return wc
}

func parseCardNumber(s string) int {
	fields := strings.Fields(s)
	n, err := strconv.Atoi(fields[1])
	if err != nil {
		panic("failed to convert string to integer")
	}
	return n
}

func parseNumbers(s string) map[int]struct{} {
	fields := strings.Fields(s)
	parsed := make(map[int]struct{}, len(fields))
	for _, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			panic("failed to convert string to integer")
		}
		parsed[n] = struct{}{}
	}
	return parsed
}

func parseCard(line string) Card {
	var chunks []string
	chunks = strings.Split(line, ":")
	cardNumber := parseCardNumber(chunks[0])
	chunks = strings.Split(chunks[1], "|")
	return Card{
		N:       cardNumber,
		Winning: parseNumbers(chunks[0]),
		My:      parseNumbers(chunks[1]),
	}
}

func solve1Line(line string) int {
	card := parseCard(line)
	wc := card.getWinCount()
	if wc == 0 {
		return 0
	}
	return 1 << (wc - 1)
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
