package day08

import (
	"fmt"
	"regexp"
	"strings"
)

var NodeLineRe = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

type Node string

type Choice struct {
	Left  Node
	Right Node
}

type Tree map[Node]Choice

type Direction int

const (
	Left Direction = iota
	Right
)

type Path struct {
	Index int
	Path  []Direction
}

func (p *Path) NextDirection() Direction {
	if p.Index == len(p.Path) {
		p.Index = 0
	}
	var next = p.Path[p.Index]
	p.Index++
	return next
}

func parsePath(in string) Path {
	var path = make([]Direction, len(in))
	for i, c := range in {
		switch c {
		case 'L':
			path[i] = Left
		case 'R':
			path[i] = Right
		default:
			panic(fmt.Sprintf("unknown direction: %v", c))
		}
	}
	return Path{0, path}
}

func parseTree(in string) Tree {
	var tree = make(Tree)
	for _, line := range strings.Split(in, "\n") {
		var matches = NodeLineRe.FindStringSubmatch(line)
		var choice = Choice{
			Left:  Node(matches[2]),
			Right: Node(matches[3]),
		}
		var node = Node(matches[1])
		tree[node] = choice
	}
	return tree
}

const StartingNode = Node("AAA")
const FinalNode = Node("ZZZ")

func (s Solver) Solve1(in string) interface{} {
	var chunks = strings.Split(in, "\n\n")

	var path = parsePath(chunks[0])
	var tree = parseTree(chunks[1])

	var answer int
	var node = StartingNode
	for {
		var direction = path.NextDirection()
		if node == FinalNode {
			return answer
		}
		switch direction {
		case Left:
			node = tree[node].Left
		case Right:
			node = tree[node].Right
		}
		answer++
	}
}
