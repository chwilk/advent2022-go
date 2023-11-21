package day02

import (
	"advent2022-go/helpers"
	"strings"
	"testing"
)

const example = `A Y
B X
C Z`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := 15
		got := Day02{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.Check(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := 12
		got := Day02{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.Check(t, expect, got)
	})
}