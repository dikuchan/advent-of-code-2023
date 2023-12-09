package day09

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

func isArrayOfZeroes(array []int) bool {
	for _, value := range array {
		if value != 0 {
			return false
		}
	}
	return true
}

func getNextSequence(sequence []int) []int {
	var nextSequence = make([]int, len(sequence)-1)
	for i := 0; i < len(sequence)-1; i++ {
		nextSequence[i] = sequence[i+1] - sequence[i]
	}
	return nextSequence
}

func solve1Line(in string) int {
	var sequence = utils.AtoiArray(strings.Fields(in))
	var sequenceEnd = make([]int, 0)
	var answer int
	for {
		sequenceEnd = append(sequenceEnd, sequence[len(sequence)-1])
		if isArrayOfZeroes(sequence) {
			return utils.Sum(sequenceEnd...)
		}
		sequence = getNextSequence(sequence)
		answer++
	}
}

func (s Solver) Solve1(in string) interface{} {
	var answer int
	for _, line := range strings.Split(in, "\n") {
		answer += solve1Line(line)
	}
	return answer
}
