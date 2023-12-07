package day06

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

type Race struct {
	Time   int
	Record int
}

func NewRace(time int, distance int) Race {
	return Race{time, distance}
}

func (r Race) CountOptions() int {
	var options int
	for speed := 0; speed < r.Time; speed++ {
		var distance = (r.Time - speed) * speed
		if distance > r.Record {
			options++
		}
	}
	return options
}

type Leaderboard struct {
	Races []Race
}

func parseLeaderboard(in string) Leaderboard {
	var chunks = strings.Split(in, "\n")

	var timeLine = strings.Split(chunks[0], ":")
	var time = utils.AtoiArray(strings.Fields(timeLine[1]))

	var distanceLine = strings.Split(chunks[1], ":")
	var distance = utils.AtoiArray(strings.Fields(distanceLine[1]))

	var races = make([]Race, len(distance))
	for i := 0; i < len(races); i++ {
		races[i] = NewRace(time[i], distance[i])
	}
	return Leaderboard{races}
}

func solve(leaderboard Leaderboard) int {
	var answer = 1
	for _, race := range leaderboard.Races {
		answer *= race.CountOptions()
	}
	return answer
}

func (s Solver) Solve1(in string) interface{} {
	return solve(parseLeaderboard(in))
}
