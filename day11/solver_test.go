package day11

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed data/in_test.txt
var in string

const (
	answer = 374
)

var solver = Solver{}

func TestSolve1(t *testing.T) {
	assert.Equal(t, answer, solver.Solve1(in).(int))
}
