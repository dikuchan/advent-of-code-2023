package day02

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed data/in_test.txt
var in string

const (
	answer01 = 8
	answer02 = 2286
)

var solver = Solver{}

func TestSolve1(t *testing.T) {
	assert.Equal(t, answer01, solver.Solve1(in).(int))
}

func TestSolve2(t *testing.T) {
	assert.Equal(t, answer02, solver.Solve2(in).(int))
}
