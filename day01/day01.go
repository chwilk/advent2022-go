package day01

import (
	"bufio"
	"io"
	"strconv"
)

type Day01 struct {
}

func (f Day01) Run(input io.Reader) int {
	return Calories(input)
}

func Calories(lines io.Reader) (highest int) {
	var tempSum int
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			tempSum += i
		} else {
			if tempSum > highest {
				highest = tempSum
			}
			tempSum = 0
		}
	}
	return
}
