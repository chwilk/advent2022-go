package day03

import (
	"advent2022-go/helpers"
	"strings"
	"testing"
)

const example string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "157"
		got := Day03{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "70"
		got := Day03{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestPrioritize(t *testing.T) {
	t.Run("Test prioritize function adjusts alpha", func(t *testing.T) {
		array := []byte{'a', 'b', 'c', 'A', 'B', 'C'}
		expect := []byte{1, 2, 3, 27, 28, 29}
		helpers.CheckBytes(t, expect, prioritize(array))
	})
}
