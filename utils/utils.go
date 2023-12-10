package utils

import (
	"strconv"

	cs "golang.org/x/exp/constraints"
)

func RuneToInteger(r rune) int {
	return int(r - '0')
}

func IsDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func Atoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic("failed to convert string to integer")
	}
	return value
}

func AtoiArray(sarray []string) []int {
	var array = make([]int, len(sarray))
	for i, field := range sarray {
		array[i] = Atoi(field)
	}
	return array
}

func FindMin[T cs.Integer](array []T) T {
	var r = array[0]
	for _, value := range array {
		if value < r {
			r = value
		}
	}
	return r
}

func FindMax[T cs.Integer](array []T) T {
	var r = array[0]
	for _, value := range array {
		if value > r {
			r = value
		}
	}
	return r
}

func Max[T cs.Integer](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func GCD[T cs.Integer](x, y T) T {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func LCM[T cs.Integer](values ...T) T {
	if len(values) == 1 {
		return values[0]
	}
	var x = values[0]
	var y = values[1]
	var r = x * y / GCD(x, y)
	values = values[2:]
	for i := 0; i < len(values); i++ {
		r = LCM(r, values[i])
	}
	return r
}

func Sum[T cs.Integer](values ...T) T {
	var r T
	for _, value := range values {
		r += value
	}
	return r
}
