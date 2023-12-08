package day08

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

func (n Node) IsStarting() bool {
	return n[len(n)-1] == 'A'
}

func (n Node) IsFinal() bool {
	return n[len(n)-1] == 'Z'
}

func (t Tree) GetStartingNodes() []Node {
	var nodes = make([]Node, 0)
	for node := range t {
		if node.IsStarting() {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func solveForNode(tree Tree, path Path, node Node) int64 {
	var steps int64
	for {
		var direction = path.NextDirection()
		if node.IsFinal() {
			return steps
		}
		var nextNode Node
		switch direction {
		case Left:
			nextNode = tree[node].Left
		case Right:
			nextNode = tree[node].Right
		}
		node = nextNode
		steps++
	}
}

func (s Solver) Solve2(in string) interface{} {
	var chunks = strings.Split(in, "\n\n")

	var path = parsePath(chunks[0])
	var tree = parseTree(chunks[1])

	var nodes = tree.GetStartingNodes()
	var steps = make([]int64, len(nodes))
	for i, node := range nodes {
		steps[i] = solveForNode(tree, path, node)
	}
	return utils.LCM(steps...)
}
