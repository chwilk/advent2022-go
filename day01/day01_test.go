package day01

import (
	"advent2022-go/helpers"
	"strings"
	"testing"
)

const example string = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestCalories(t *testing.T) {
	t.Run("Calories part A", func(t *testing.T) {
		expect := 24000
		got := Calories(strings.NewReader(example))[0]

		helpers.Check(t, expect, got)
	})
	t.Run("Calories part B", func(t *testing.T) {
		expect := 45000
		a := Calories(strings.NewReader(example))
		got := a[0] + a[1] + a[2]

		helpers.Check(t, expect, got)
	})
}

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := 24000
		got := Day01{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.Check(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := 45000
		got := Day01{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.Check(t, expect, got)
	})
}
