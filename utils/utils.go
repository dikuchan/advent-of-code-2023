package utils

import "strconv"

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

func FindMin(array []int) int {
	var answer = array[0]
	for _, value := range array {
		if value < answer {
			answer = value
		}
	}
	return answer
}

func FindMax(array []int) int {
	var answer = array[0]
	for _, value := range array {
		if value > answer {
			answer = value
		}
	}
	return answer
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
