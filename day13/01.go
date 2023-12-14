package day13

import (
	"strings"
)

func isMirrorH(lines []string, i, j, N int) bool {
	for k := 0; i-k > 0 && i+k < N; k++ {
		if lines[j][i-k-1] != lines[j][i+k] {
			return false
		}
	}
	return true
}

func isMirrorV(lines []string, i, j, N int) bool {
	for k := 0; j-k > 0 && j+k < N; k++ {
		if lines[j-k-1][i] != lines[j+k][i] {
			return false
		}
	}
	return true
}

func solveNote(note string) int {
	var lines = strings.Split(note, "\n")

	var Y = len(lines)
	var X = len(lines[0])

	for i := 1; i < X; i++ {
		var mirror = true
		for j := 0; j < Y; j++ {
			if !isMirrorH(lines, i, j, X) {
				mirror = false
				break
			}
		}
		if mirror {
			return i
		}
	}
	for j := 1; j < Y; j++ {
		var mirror = true
		for i := 0; i < X; i++ {
			if !isMirrorV(lines, i, j, Y) {
				mirror = false
				break
			}
		}
		if mirror {
			return 100 * j
		}
	}
	return 0
}

func (Solver) Solve1(in string) interface{} {
	var answer int
	for _, note := range strings.Split(in, "\n\n") {
		answer += solveNote(note)
	}
	return answer
}
