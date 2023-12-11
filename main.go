package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/dikuchan/aoc-2023/day01"
	"github.com/dikuchan/aoc-2023/day02"
	"github.com/dikuchan/aoc-2023/day03"
	"github.com/dikuchan/aoc-2023/day04"
	"github.com/dikuchan/aoc-2023/day05"
	"github.com/dikuchan/aoc-2023/day06"
	"github.com/dikuchan/aoc-2023/day07"
	"github.com/dikuchan/aoc-2023/day08"
	"github.com/dikuchan/aoc-2023/day09"
	"github.com/dikuchan/aoc-2023/day10"
	"github.com/dikuchan/aoc-2023/day11"
)

type Solver interface {
	Solve1(in string) interface{}
	Solve2(in string) interface{}
}

//go:embed day01/data/in.txt
var in01 string

//go:embed day02/data/in.txt
var in02 string

//go:embed day03/data/in.txt
var in03 string

//go:embed day04/data/in.txt
var in04 string

//go:embed day05/data/in.txt
var in05 string

//go:embed day06/data/in.txt
var in06 string

//go:embed day07/data/in.txt
var in07 string

//go:embed day08/data/in.txt
var in08 string

//go:embed day09/data/in.txt
var in09 string

//go:embed day10/data/in.txt
var in10 string

//go:embed day11/data/in.txt
var in11 string

var Solvers = map[string]Solver{
	"01": day01.Solver{},
	"02": day02.Solver{},
	"03": day03.Solver{},
	"04": day04.Solver{},
	"05": day05.Solver{},
	"06": day06.Solver{},
	"07": day07.Solver{},
	"08": day08.Solver{},
	"09": day09.Solver{},
	"10": day10.Solver{},
	"11": day11.Solver{},
}
var In = map[string]string{
	"01": in01,
	"02": in02,
	"03": in03,
	"04": in04,
	"05": in05,
	"06": in06,
	"07": in07,
	"08": in08,
	"09": in09,
	"10": in10,
	"11": in11,
}

var task string
var part int

func main() {
	flag.StringVar(&task, "task", "", "")
	flag.IntVar(&part, "part", 0, "")

	flag.Parse()

	if solver, ok := Solvers[task]; ok {
		var answer interface{}
		switch part {
		case 1:
			answer = solver.Solve1(In[task])
		case 2:
			answer = solver.Solve2(In[task])
		default:
			panic(fmt.Sprintf("No such part: %v", part))
		}
		fmt.Println(answer)
	} else {
		panic(fmt.Sprintf("No such task: %v", task))
	}
}
