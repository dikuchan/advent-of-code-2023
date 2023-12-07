package day02

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

func solve2Line(line string) int {
	var chunks = strings.Split(line, ":")
	var cubeSets = strings.Split(chunks[1], ";")

	var counter = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, cubeSet := range cubeSets {
		var cubes = strings.Split(cubeSet, ",")
		for color, n := range countCubes(cubes) {
			counter[color] = utils.Max(counter[color], n)
		}
	}

	var answer = 1
	for _, value := range counter {
		answer *= value
	}
	return answer
}

func (s Solver) Solve2(in string) interface{} {
	return solve(in, solve2Line)
}
