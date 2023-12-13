package day12

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

type Record struct {
	springs string
	sizes   []int
	size    int
}

func NewRecord(springs string, sizes []int, size int) Record {
	return Record{springs, sizes, size}
}

func parseRecords(in string) []Record {
	var lines = strings.Split(in, "\n")
	var records = make([]Record, len(lines))
	for i, line := range lines {
		var chunks = strings.Split(line, " ")
		var springs = chunks[0]
		var sizes = utils.AtoiArray(strings.Split(chunks[1], ","))
		records[i] = Record{springs, sizes, 0}
	}
	return records
}

type Cache map[string]map[int]map[int]int

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Get(r Record) int {
	var sizesHash = utils.Sum(r.sizes...)
	sizesMap, ok := (*c)[r.springs]
	if !ok {
		return -1
	}
	sizeMap, ok := sizesMap[sizesHash]
	if !ok {
		return -1
	}
	answer, ok := sizeMap[r.size]
	if !ok {
		return -1
	}
	return answer
}

func (c *Cache) Set(r Record, value int) {
	var sizesHash = utils.Sum(r.sizes...)
	sizesMap, ok := (*c)[r.springs]
	if !ok {
		(*c)[r.springs] = make(map[int]map[int]int)
	}
	if _, ok := sizesMap[sizesHash]; !ok {
		(*c)[r.springs][sizesHash] = make(map[int]int)
	}
	(*c)[r.springs][sizesHash][r.size] = value
}

func (r Record) solveForChar(cache *Cache, c byte) int {
	if c == '#' {
		return NewRecord(r.springs[1:], r.sizes, r.size+1).solve(cache)
	} else {
		if r.size > 0 {
			if len(r.sizes) > 0 && r.sizes[0] == r.size {
				return NewRecord(r.springs[1:], r.sizes[1:], 0).solve(cache)
			}
			return 0
		} else {
			return NewRecord(r.springs[1:], r.sizes, 0).solve(cache)
		}
	}
}

func (r Record) solve(cache *Cache) int {
	if len(r.springs) == 0 {
		if r.size == 0 && len(r.sizes) == 0 {
			return 1
		}
		return 0
	}
	var answer = cache.Get(r)
	if answer != -1 {
		return answer
	}
	if r.springs[0] == '?' {
		answer = r.solveForChar(cache, '.') + r.solveForChar(cache, '#')
	} else {
		answer = r.solveForChar(cache, r.springs[0])
	}
	cache.Set(r, answer)
	return answer
}

func (Solver) Solve1(in string) interface{} {
	var answer = 0
	for _, r := range parseRecords(in) {
		var cache = NewCache()
		answer += NewRecord(r.springs+".", r.sizes, r.size).solve(cache)
	}
	return answer
}
