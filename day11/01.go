package day11

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

type galaxy struct {
	x int
	y int
}

func newGalaxy(x int, y int) galaxy {
	return galaxy{x, y}
}

type image map[int]map[int]byte

func parseImage(in string) image {
	var img = make(image)
	for y, line := range strings.Split(in, "\n") {
		for x, c := range line {
			if _, ok := img[x]; !ok {
				img[x] = make(map[int]byte)
			}
			img[y][x] = byte(c)
		}
	}
	return img
}

func findDistance(p, q int, expansionMap []bool, expandedBy int) int {
	var d = utils.Abs(p - q)
	for i := utils.Min(p, q); i < utils.Max(p, q); i++ {
		if expansionMap[i] {
			d += expandedBy
		}
	}
	return d
}

func solve(in string, expandedBy int) int {
	var img = parseImage(in)

	var nX = len(img)
	var nY = len(img[0])

	var galaxies = make([]galaxy, 0)
	for y, row := range img {
		for x, c := range row {
			if c == '#' {
				galaxies = append(galaxies, newGalaxy(x, y))
			}
		}
	}

	var expansionY = make([]bool, nY)
	for y, row := range img {
		var expanded = true
		for _, c := range row {
			if c == '#' {
				expanded = false
				break
			}
		}
		expansionY[y] = expanded
	}

	var expansionX = make([]bool, nX)
	for x := 0; x < nX; x++ {
		var expanded = true
		for y := 0; y < nY; y++ {
			if img[y][x] == '#' {
				expanded = false
				break
			}
		}
		expansionX[x] = expanded
	}

	var answer int
	for i, p := range galaxies {
		for _, q := range galaxies[i+1:] {
			var dx = findDistance(p.x, q.x, expansionX, expandedBy)
			var dy = findDistance(p.y, q.y, expansionY, expandedBy)
			answer += dx + dy
		}
	}
	return answer
}

func (Solver) Solve1(in string) interface{} {
	return solve(in, 1)
}
