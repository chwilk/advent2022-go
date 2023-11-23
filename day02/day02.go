package day02

import (
	"advent2022-go/helpers"
	"bufio"
	"fmt"
	"io"
)

type Day02 struct {
}

func (f Day02) Run(input io.Reader, part int) (result string) {
	switch part {
	case helpers.PartA:
		result = fmt.Sprint(ScoreA(input))
	default:
		result = fmt.Sprint(ScoreB(input))
	}
	return
}

func parseLine(round []byte) (elf, self int) {
	elf = int(round[0] - 65)
	self = int(round[2] - 88)
	return
}
func ScoreA(input io.Reader) (score int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		elf, self := parseLine(scanner.Bytes())
		score += self + 1 // score for your throw
		switch self - elf {
		case 1:
			fallthrough
		case -2:
			score += 6
		case 0:
			score += 3
		default:
		}
	}
	return
}

func ScoreB(input io.Reader) (score int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		elf, outcome := parseLine(scanner.Bytes())
		score += 3 * (outcome)
		switch outcome {
		case 0:
			score += (elf+2)%3 + 1
		case 1:
			score += elf + 1
		case 2:
			score += (elf+1)%3 + 1
		}
	}
	return
}
