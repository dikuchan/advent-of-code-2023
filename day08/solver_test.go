package day08

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed data/in_test01.txt
var in01 string

//go:embed data/in_test02.txt
var in02 string

//go:embed data/in_test03.txt
var in03 string

const (
	answer01       = 2
	answer02       = 6
	answer03 int64 = 6
)

var solver = Solver{}

func TestSolve1_1(t *testing.T) {
	assert.Equal(t, answer01, solver.Solve1(in01).(int))
}

func TestSolve1_2(t *testing.T) {
	assert.Equal(t, answer02, solver.Solve1(in02).(int))
}

func TestSolve2(t *testing.T) {
	assert.Equal(t, answer03, solver.Solve2(in03).(int64))
}
