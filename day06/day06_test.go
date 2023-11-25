package day06

import (
	"advent2022-go/helpers"
	"strings"
	"testing"
)

const example string = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "7"
		got := Day06{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "19"
		got := Day06{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestSum(t *testing.T) {
	data := []int{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1}
	got := sum(data)
	expect := 4
	helpers.Check(t, expect, got)
}