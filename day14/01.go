package day14

import (
	"bytes"
	"strings"
)

func (Solver) Solve1(in string) interface{} {
	var lines = strings.Split(in, "\n")

	var Y = len(lines)
	var X = len(lines[0])

	var data = make([][]byte, X)
	for x := 0; x < X; x++ {
		data[x] = make([]byte, Y)
	}
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			data[x][y] = byte(lines[y][x])
		}
	}

	for _, r := range data {
		for _, s := range bytes.Split(r, []byte{'#'}) {
			var n = bytes.Count(s, []byte{'O'})
			for i := 0; i < len(s); i++ {
				if i < n {
					s[i] = 'O'
				} else {
					s[i] = '.'
				}
			}
		}
	}
	var r = 0
	for _, s := range data {
		for i := 0; i < len(s); i++ {
			if s[i] == 'O' {
				r += M - i
			}
		}
	}
	return r
}
