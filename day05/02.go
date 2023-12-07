package day05

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

type SeedRange struct {
	Begin  int
	Length int
	End    int
}

func NewSeedRange(begin int, length int) SeedRange {
	return SeedRange{begin, length, begin + length}
}

func parseSeedRanges(s string) []SeedRange {
	var chunks = strings.Split(s, ":")
	var fields = strings.Fields(chunks[1])

	var seedRanges = make([]SeedRange, 0)
	for i := 0; i < len(fields); i += 2 {
		var begin = utils.Atoi(fields[i])
		var length = utils.Atoi(fields[i+1])
		seedRanges = append(seedRanges, NewSeedRange(begin, length))
	}
	return seedRanges
}

func (s Solver) Solve2(in string) interface{} {
	var chunks = strings.Split(in, "\n\n")
	var seedRanges = parseSeedRanges(chunks[0])
	var almanac = parseAlmanac(chunks[1:])
	var locationCache = make(map[int]int)

	var answer = 2 << 31
	for _, seedRange := range seedRanges {
		for seed := seedRange.Begin; seed < seedRange.End; seed++ {
			location, ok := locationCache[seed]
			if !ok {
				location = almanac.FindLocation(seed)
			}
			if location < answer {
				answer = location
			}
		}
	}
	return answer
}
