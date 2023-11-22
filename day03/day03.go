package day03

import (
	"advent2022-go/helpers"
	"bufio"
	"io"
)

type Day03 struct {
}

func (f Day03) Run(input io.Reader, part int) (result int) {
	switch part {
	case helpers.PartA:
		result = PartA(input)
	default:
		
	}
	return
}

//var items [53]bool

func PartA(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
	    var tmpItems [53]bool
		contents := prioritize(scanner.Bytes())
		compartment := len(contents)/2
		for _,v := range contents[:compartment-1] {
			tmpItems[v] = true
		}
		for _,v := range contents[compartment:] {
			if tmpItems[v] {
				answer += int(v)
				break
			}
		}
	}
	return
}

func prioritize(array []byte) []byte {
	for i, v := range array {
		if  v > 96 {
			array[i] -= 96
		} else {
			array[i] -= 38
		}
	}
	return array
}