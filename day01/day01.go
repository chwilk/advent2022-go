package day01

import (
	"advent2022-go/helpers"
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type Day01 struct {
}

func (f Day01) Run(input io.Reader, part int) (result string) {
	switch part {
	case helpers.PartA:
		result = fmt.Sprint(Calories(input)[0])
	default:
		a := Calories(input)
		result = fmt.Sprint(a[0] + a[1] + a[2])
	}
	return
}

func Calories(lines io.Reader) (highest []int) {
	var tempSum int
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err == nil {
			tempSum += i
		} else {
			highest = append(highest, tempSum)
			tempSum = 0
		}
	}
	highest = append(highest, tempSum)
	sort.Slice(highest, func(i, j int) bool { return highest[j] < highest[i] })
	return
}
