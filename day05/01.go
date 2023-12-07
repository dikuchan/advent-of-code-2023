package day05

import (
	"regexp"
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

var MappingNameRe = regexp.MustCompile(`(\w+)-to-(\w+)`)

type Range struct {
	Destination int
	Source      int
	Length      int
}

func NewRange(dst int, src int, length int) Range {
	return Range{dst, src, length}
}

type Mapping struct {
	Source      string
	Destination string
	Ranges      []Range
}

func (m Mapping) Map(value int) int {
	for _, rng := range m.Ranges {
		if value < rng.Source || value >= rng.Source+rng.Length {
			continue
		}
		return rng.Destination + (value - rng.Source)
	}
	return value
}

type Almanac struct {
	Mappings map[string]Mapping
}

func (a Almanac) FindLocation(seed int) int {
	var source = "seed"
	var answer = seed
	for {
		if mapping, ok := a.Mappings[source]; ok {
			answer = mapping.Map(answer)
			source = mapping.Destination
		} else {
			return answer
		}
	}
}

func parseSeeds(s string) []int {
	var chunks = strings.Split(s, ":")
	var seeds = strings.Fields(chunks[1])
	return utils.AtoiArray(seeds)
}

func parseMappingName(s string) (string, string) {
	var matches = MappingNameRe.FindStringSubmatch(s)
	return matches[1], matches[2]
}

func parseMapping(s string) Mapping {
	var chunks = strings.Split(s, "\n")
	source, destination := parseMappingName(chunks[0])

	var ranges = make([]Range, len(chunks)-1)
	for i, chunk := range chunks[1:] {
		var fields = strings.Fields(chunk)
		var values = utils.AtoiArray(fields)
		ranges[i] = NewRange(values[0], values[1], values[2])
	}

	return Mapping{source, destination, ranges}
}

func parseAlmanac(lines []string) Almanac {
	var mappings = make(map[string]Mapping)
	for _, line := range lines {
		mapping := parseMapping(line)
		mappings[mapping.Source] = mapping
	}
	return Almanac{mappings}
}

func (s Solver) Solve1(in string) interface{} {
	var chunks = strings.Split(in, "\n\n")
	var seeds = parseSeeds(chunks[0])
	var almanac = parseAlmanac(chunks[1:])

	var locations = make([]int, len(seeds))
	for i, seed := range seeds {
		locations[i] = almanac.FindLocation(seed)
	}
	return utils.FindMin(locations)
}
