package day02

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

var Cubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func parseGameID(s string) int {
	var id, _ = strings.CutPrefix(s, "Game ")
	return utils.Atoi(id)
}

func parseCube(cube string) (string, int) {
	cube = strings.TrimSpace(cube)
	var chunks = strings.Split(cube, " ")
	var n = utils.Atoi(chunks[0])
	var color = chunks[1]
	return color, n
}

func countCubes(cubes []string) map[string]int {
	var counter = make(map[string]int)
	for _, cube := range cubes {
		var color, n = parseCube(cube)
		counter[color] += n
	}
	return counter
}

func checkCubeSet(s string) bool {
	var cubes = strings.Split(s, ",")
	var counter = countCubes(cubes)
	for color, n := range Cubes {
		if counter[color] > n {
			return false
		}
	}
	return true
}

func solve1Line(line string) int {
	var chunks = strings.Split(line, ":")
	var cubeSets = strings.Split(chunks[1], ";")
	for _, cubeSet := range cubeSets {
		if !checkCubeSet(cubeSet) {
			return 0
		}
	}
	return parseGameID(chunks[0])
}

func solve(in string, solveLine func(string) int) int {
	var lines = strings.Split(in, "\n")
	var answer int
	for _, line := range lines {
		answer += solveLine(line)
	}
	return answer
}

func (s Solver) Solve1(in string) interface{} {
	return solve(in, solve1Line)
}
