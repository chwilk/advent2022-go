package main

import (
	"advent2022-go/day01"
	"advent2022-go/helpers"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Define an interface instantiated by each day
type Days interface {
	Run(io.Reader, int) int
}

// Run against day of your choosing, default all
func main() {
	args := os.Args[1:]
	var day int
	var err error
	if len(args) >= 1 {
		day, err = strconv.Atoi(args[0])
		if err != nil || day > 25 || day < 1{
			usage()
		}
	}
	fmt.Printf("Running for day%02d\n", day)
	if day > 0 {
		runDay(getDay(day), day, helpers.PartA)
	} else {
		runDay(getDay(1), 1, helpers.PartA)
	}
}

func usage() {
	fmt.Printf("Usage:\n\t%v [day]\n", os.Args[0])
	os.Exit(1)
}

func getDay(day int) Days {
	switch day {
	default:
		return day01.Day01{}
	}
}

func runDay(d Days, day, part int) {
	filename := fmt.Sprintf("problems/input%02d.dat", day)
	data, err := os.ReadFile(filename)
	if err == nil {
		answerA := d.Run(bytes.NewReader(data), helpers.PartA)
		answerB := d.Run(bytes.NewReader(data), helpers.PartB)
		fmt.Printf("Day %2d answers are %d and %d\n",day,answerA, answerB)
	}
}