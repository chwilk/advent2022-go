package day03

import (
	"advent2022-go/helpers"
	"bufio"
	"fmt"
	"io"
)

type Day03 struct {
}

func (f Day03) Run(input io.Reader, part int) (result string) {
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

func PartB(lines io.Reader) (answer int) {
	scanner := bufio.NewScanner(lines)
	var items [53]bool
	var groupMember int
	for scanner.Scan() {
	    var tmpItems [53]bool
		contents := prioritize(scanner.Bytes())
		for _,v := range contents {
			tmpItems[v] = true
		}
		switch groupMember {
		case 0:
			for i := range items {
				items[i] = tmpItems[i]
			}
		case 1:
			for i := range items {
				items[i] = items[i] && tmpItems[i]
			}
		case 2:
			for i := range items {
				if items[i] && tmpItems[i] {
					answer += i
				}
				items[i] = false
			}
			groupMember = -1
		}
		groupMember ++
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
