package day05

import (
	"advent2022-go/helpers"
	"reflect"
	"strings"
	"testing"
)

const example string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestRun(t *testing.T) {
	t.Run("Run function part A", func(t *testing.T) {
		expect := "CMZ"
		got := Day05{}.Run(strings.NewReader(example), helpers.PartA)
		helpers.CheckString(t, expect, got)
	})
	t.Run("Run function part B", func(t *testing.T) {
		expect := ""
		got := Day05{}.Run(strings.NewReader(example), helpers.PartB)
		helpers.CheckString(t, expect, got)
	})
}

func TestStack(t *testing.T) {
	s := stack{}
	s = s.Push('A')
	s = s.Push('B')
	s = s.Push('C')
	s = s.Insert('D')
	expect := byte('C')
	_, got := s.Pop()
	if got != expect {
		t.Errorf("Expected %d, got %d", expect, got)
	}
}

func TestMove(t *testing.T) {
	expect := [9]stack{{}, {'B', 'A'}, {'C'}}
	stacks := [9]stack{}
	stacks[0] = stacks[0].Push('A')
	stacks[1] = stacks[1].Push('B')
	stacks[2] = stacks[2].Push('C')
	stacks = move(stacks, 1, 1, 2)
	if !reflect.DeepEqual(expect, stacks) {
		t.Errorf("Expected %d, got %d", expect, stacks)
	}
}
