package day01

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed data/in_test01.txt
var in01 string

//go:embed data/in_test02.txt
var in02 string

const (
	answer01 = 142
	answer02 = 281
)

var solver = Solver{}

func TestSolve1(t *testing.T) {
	assert.Equal(t, answer01, solver.Solve1(in01).(int))
}

func TestSolve2(t *testing.T) {
	assert.Equal(t, answer02, solver.Solve2(in02).(int))
}
