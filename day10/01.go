package day10

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

type Vertex struct {
	x int
	y int
	v byte
}

func NewVertex(x int, y int, v byte) Vertex {
	return Vertex{x, y, v}
}

type Graph map[Vertex][]Vertex

func parseGraph(in string) Graph {
	var graph = make(map[Vertex][]Vertex)

	var lines = strings.Split(in, "\n")
	var MY = len(lines)
	var MX = len(lines[0])

	for y, line := range strings.Split(in, "\n") {
		for x, c := range line {
			var b = byte(c)
			if b == '.' {
				continue
			}
			var v = NewVertex(x, y, b)
			var ws = make([]Vertex, 0)
			for _, w := range []struct {
				dx int
				dy int
				ps []byte
			}{
				{-1, 0, []byte{'-', 'L', 'F', 'S'}},
				{1, 0, []byte{'-', 'J', '7', 'S'}},
				{0, -1, []byte{'|', '7', 'F', 'S'}},
				{0, 1, []byte{'|', 'L', 'J', 'S'}},
			} {
				var nx = x + w.dx
				var ny = y + w.dy
				if nx < 0 || nx >= MX || ny < 0 || ny >= MY {
					continue
				}
				for _, p := range w.ps {
					if lines[ny][nx] == p {
						ws = append(ws, NewVertex(nx, ny, p))
						continue
					}
				}
			}
			graph[v] = ws
		}
	}
	return graph
}

func (g Graph) findStartingVertex() Vertex {
	for n := range g {
		if n.v == 'S' {
			return n
		}
	}
	panic("no starting vertex found")
}

func (Solver) Solve1(in string) interface{} {
	var graph = parseGraph(in)
	var sv = graph.findStartingVertex()

	var sizes = make(map[Vertex]int)

	var visited = make(map[Vertex]bool)
	visited[sv] = true

	var queue = utils.NewQueue[Vertex]()
	queue.Push(sv)

	for !queue.IsEmpty() {
		var v = queue.Pop()
		for _, w := range graph[v] {
			if !visited[w] {
				visited[w] = true
				queue.Push(w)
			}
			sizes[w] = sizes[v] + 1
		}
	}
	var answer int
	for _, s := range sizes {
		answer = utils.Max(answer, s)
	}
	return answer - 1
}
