package day02

import (
	"advent2022-go/helpers"
	"io"
)

type Day02 struct {
}

func (f Day02) Run(input io.Reader, part int) (result int) {
	switch part {
	case helpers.PartA:
		result = ScoreA(input)
	default:
		result = ScoreB(input)
	}
	return
}

func ScoreA(input io.Reader) (score int) {
	return
}

func ScoreB(input io.Reader) (score int) {
	return
}