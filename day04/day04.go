package day04

import (
	"advent2022-go/helpers"
	"bufio"
	"io"
	"regexp"
	"strconv"
)

type Day04 struct {
}

func (f Day04) Run(input io.Reader, part int) (result int) {
	switch part {
	case helpers.PartA:
		result = PartA(input)
	default:
		result = PartB(input)
	}
	return
}

func PartA(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		a, b := parseLine(scanner.Text())
		if contains(a, b) {
			answer ++
		}
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
		a, b := parseLine(scanner.Text())
		if overlaps(a, b) {
			answer ++
		}
	}
	return
}

func contains(a, b [2]int) bool {
	switch {
	case a[0] <= b[0] && a[1] >= b[1]:
		fallthrough
	case a[0] >= b[0] && a[1] <= b[1]:
		return true
	default:
		return false
	}
}

func overlaps(a, b [2]int) bool {
	if contains(a,b) { return true}
	switch {
	case a[0] <= b[0] && a[1] >= b[0]:
		fallthrough
	case b[0] <= a[0] && b[1] >= a[0]:
		return true
	default:
		return false
	}
}

func parseLine(s string) (a, b [2]int) {
    tokens := regexp.MustCompile("[,-]").Split(s, 4)
	var err error
	a[0], err = strconv.Atoi(tokens[0])
	if err != nil {panic(err)}
	a[1], err = strconv.Atoi(tokens[1])
	if err != nil {panic(err)}
	b[0], err = strconv.Atoi(tokens[2])
	if err != nil {panic(err)}
	b[1], err = strconv.Atoi(tokens[3])
	if err != nil {panic(err)}
	return
}