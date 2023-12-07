package day06

import (
	"strings"

	"github.com/dikuchan/aoc-2023/utils"
)

func parseLeaderboard2(in string) Leaderboard {
	var chunks = strings.Split(in, "\n")

	var timeLine = strings.Split(chunks[0], ":")
	var time = utils.Atoi(strings.ReplaceAll(timeLine[1], " ", ""))

	var distanceLine = strings.Split(chunks[1], ":")
	var distance = utils.Atoi(strings.ReplaceAll(distanceLine[1], " ", ""))

	var races = []Race{NewRace(time, distance)}
	return Leaderboard{races}
}

func (s Solver) Solve2(in string) interface{} {
	return solve(parseLeaderboard2(in))
}
