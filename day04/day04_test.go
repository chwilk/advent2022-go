package day04

import (
	"advent2022-go/helpers"
	"strings"
	"testing"
)

const example string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "2"
		got := Day04{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := "4"
		got := Day04{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestContains(t *testing.T) {
	t.Run("First contains second", func(t *testing.T) {
		a := [2]int{2, 8}
		b := [2]int{3, 7}
		expect := true
		got := contains(a, b)
		if expect != got {
			t.Errorf("Wrong result for %q, %q", a, b)
		}
	})
	t.Run("Second contains first", func(t *testing.T) {
		a := [2]int{3, 7}
		b := [2]int{2, 8}
		expect := true
		got := contains(a, b)
		if expect != got {
			t.Errorf("Wrong result for %q, %q", a, b)
		}
	})
}

func TestOverlaps(t *testing.T) {
	t.Run("Simple overlap", func(t *testing.T) {
		a := [2]int{5, 7}
		b := [2]int{7, 9}
		expect := true
		got := overlaps(a, b)
		if expect != got {
			t.Errorf("Wrong result for %v %v", a, b)
		}
		got = overlaps(b, a)
		if expect != got {
			t.Errorf("Wrong result for %v %v", b, a)
		}
	})
	t.Run("Single value overlap", func(t *testing.T) {
		a := [2]int{6, 6}
		b := [2]int{4, 6}
		expect := true
		got := overlaps(b, a)
		if expect != got {
			t.Errorf("Wrong result for %v %v", a, b)
		}
	})
}

func TestParseLine(t *testing.T) {
	s := "12-148,23-367"
	a, b := parseLine(s)
	expect := [4]int{12, 148, 23, 367}
	for i := 0; i < 2; i++ {
		switch {
		case a[i] != expect[i]:
			fallthrough
		case b[i] != expect[i+2]:
			t.Errorf("Got %v and %v, expected %v", a, b, expect)
		default:
		}
	}
}
