package day12

import "strings"

func (s Solver) Solve2(in string) interface{} {
	var answer = 0
	for _, r := range parseRecords(in) {
		var cache = NewCache()

		var sizesN = len(r.sizes)
		var sizes = make([]int, sizesN*5)
		for i := 0; i < 5; i++ {
			for j, s := range r.sizes {
				sizes[j+(i*sizesN)] = s
			}
		}
		var springsArray = make([]string, 5)
		for i := 0; i < 5; i++ {
			springsArray[i] = r.springs
		}
		var springs = strings.Join(springsArray, "?") + "."

		answer += NewRecord(springs, sizes, r.size).solve(cache)
	}
	return answer
}
