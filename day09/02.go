package day09

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

func extrapolate(array []int) int {
	var r int
	for i := len(array) - 1; i >= 0; i-- {
		r = array[i] - r
	}
	return r
}

func solve2Line(in string) int {
	var sequence = utils.AtoiArray(strings.Fields(in))
	var sequenceBegin = make([]int, 0)
	var answer int
	for {
		sequenceBegin = append(sequenceBegin, sequence[0])
		if isArrayOfZeroes(sequence) {
			return extrapolate(sequenceBegin)
		}
		sequence = getNextSequence(sequence)
		answer++
	}
}

func (s Solver) Solve2(in string) interface{} {
	var answer int
	for _, line := range strings.Split(in, "\n") {
		answer += solve2Line(line)
	}
	return answer
}
