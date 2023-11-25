package day06

import (
	"advent2022-go/helpers"
	"bufio"
	"fmt"
	"io"
)

type Day06 struct {
}

func (f Day06) Run(input io.Reader, part int) (result string) {
	switch part {
	case helpers.PartA:
		result = fmt.Sprint(PartA(input, 4))
	default:
		result = fmt.Sprint(PartA(input, 14))
	}
	return
}

func PartA(lines io.Reader, length int) int {
	scanner := bufio.NewScanner(lines)
	for scanner.Scan() {
		buffer := scanner.Bytes()
		for i:= length-1; i< len(buffer); i++ {
			data := make([]int, 26)
			for _, v := range buffer[i-length+1:i+1] {
				data[v-97] = 1
			}
			if sum(data) == length {
				return i+1
			}
			
		}
	}
	return 0
}

func sum(data []int) (answer int) {
	for _,v := range data {
		answer += v
	}
	return
}