package day1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExample(t *testing.T) {
	exampleIn := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	d := Day1{}
	d.SetInput(exampleIn)
	result := d.SolvePart1()
	if diff := cmp.Diff("142", result); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestExamplePart2(t *testing.T) {
	exampleIn := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	d := Day1{}
	d.SetInput(exampleIn)
	result := d.SolvePart2()
	if diff := cmp.Diff("281", result); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestLineValue2(t *testing.T) {
	table := []struct {
		line string
		want int
	}{
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"9zml", 99},
		{"foofivebar", 55},
		{"eightwo", 82},
		{"19", 19},
	}

	for _, tst := range table {
		got := lineValue2(tst.line)
		if diff := cmp.Diff(tst.want, got); diff != "" {
			t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
		}
	}
}
