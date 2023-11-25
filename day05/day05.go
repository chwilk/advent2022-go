package day05

import (
	"advent2022-go/helpers"
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Day05 struct {
}

type stack []byte

func (s stack) Push(v byte) stack {
	return append(s, v)
}

func (s stack) Insert(v byte) stack {
	//tmp := make([]byte, len(s)+1)
	//tmp[0] = v
	return append([]byte{v}, s...)
}

func (s stack) Pop() (stack, byte) {
	l := len(s)
	if l > 0 {
		return s[:l-1], s[l-1]
	} else {
		return s, 0
	}
}

func (f Day05) Run(input io.Reader, part int) (result string) {
	switch part {
	case helpers.PartA:
		result = PartA(input)
	default:
		result = PartB(input)
	}
	return
}

func PartA(lines io.Reader) (answer string) {
	scanner := bufio.NewScanner(lines)
	stacks := [9]stack{}

	mapping := true
	for scanner.Scan() {
		if mapping {
			line := scanner.Bytes()
			len := len(line)
			for i := 1; i < len; i += 4 {
				switch {
				case line[i] == 32:
					// noop on space
				case line[i] < 65: // reached numbers, switch modes
					mapping = false
					break
				case line[i] < 91: // is an upper alpha
					stacks[i/4] = stacks[i/4].Insert(line[i])
				default:
					panic("Got bytes we did not expect")
				}
			}
		} else {
			words := strings.Split(scanner.Text(), " ")
			if words[0] == "move" {
				count, err := strconv.Atoi(words[1])
				if err != nil {
					panic("Failure parsing move count")
				}
				from, err := strconv.Atoi(words[3])
				if err != nil {
					panic("Failure parsing move source")
				}
				to, err := strconv.Atoi(words[5])
				if err != nil {
					panic("Failure parsing move dest")
				}
				stacks = move(stacks, count, from, to)
			}
		}
	}
	return printStack(stacks)
}

func move(stacks [9]stack, count, from, to int) [9]stack {
	var tmp byte
	from -= 1
	to -= 1
	if len(stacks[from]) >= count {
		for i := 0; i < count; i++ {
			stacks[from], tmp = stacks[from].Pop()
			stacks[to] = stacks[to].Push(tmp)
		}
	} else {
		panic("Cannot move more boxes than a stack has")
	}
	return stacks
}

func printStack(stacks [9]stack) (answer string) {
	for _,s := range stacks {
		_, tmp :=  s.Pop() 
		if tmp != 0 {
			answer += string(tmp)
		}
	}
	return
}

func PartB(lines io.Reader) (answer string) {
	scanner := bufio.NewScanner(lines)
	stacks := [9]stack{}

	for scanner.Scan() {
	}
	return printStack(stacks)
}
