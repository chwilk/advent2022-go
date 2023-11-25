package day05

import (
	"advent2022-go/helpers"
	"bufio"
	"fmt"
	"io"
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
		result = fmt.Sprint(PartA(input))
	default:
		result = fmt.Sprint(PartB(input))
	}
	return
}

func PartA(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)
	stacks := [9]stack{}

	mapping := true
	for scanner.Scan() {
		line := scanner.Bytes()

		if mapping {
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

		}
	}
	return
}

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)

	for scanner.Scan() {
	}
	return
}
