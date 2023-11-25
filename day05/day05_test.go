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
		expect := "MCD"
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
	stacks := [9]stack{{'A'}, {'B', 'D'}, {'C'}}
	t.Run("CrateMover 9000 operation", func(t *testing.T) {
		expect := [9]stack{{'A', 'D', 'B'}, nil, {'C'}}
		stacksA := move(stacks, 2, 2, 1)
		if !reflect.DeepEqual(expect, stacksA) {
			t.Errorf("Expected %d, got %d", expect, stacksA)
		}
	})
	t.Run("CrateMover 9001 operation", func(t *testing.T) {
		expect := [9]stack{{'A', 'B', 'D'}, nil, {'C'}}
		stacksB := move(stacks, 2, 2, 5)
		stacksB = move(stacksB, 2, 5, 1)
		if !reflect.DeepEqual(expect, stacksB) {
			t.Errorf("Expected %v, got %v", expect, stacksB)
		}
	})
}
