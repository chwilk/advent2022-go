package day01

import (
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
	expect := 24000
	got := Calories(strings.NewReader(example))

	if got != expect {
		t.Errorf("Expected %d got %d.", expect, got)
	}
}
