package day3

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	want := "4361"

	d2 := &Day3{}
	d2.SetInput(exampleIn)
	got := d2.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart2(t *testing.T) {
	exampleIn := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	want := "467835"

	d2 := &Day3{}
	d2.SetInput(exampleIn)
	got := d2.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestIToXY(t *testing.T) {
	cells := make([]byte, 100)
	g := grid{10, 10, cells}

	table := []struct {
		in   int
		want point
	}{
		{0, point{0, 0}},
		{5, point{5, 0}},
		{9, point{9, 0}},
		{10, point{0, 1}},
		{15, point{5, 1}},
		{19, point{9, 1}},
	}

	for _, tst := range table {
		got := g.iToXY(tst.in)
		if diff := cmp.Diff(tst.want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestXYToI(t *testing.T) {
	cells := make([]byte, 100)
	g := grid{10, 10, cells}

	table := []struct {
		in   point
		want int
	}{
		{point{0, 0}, 0},
		{point{5, 0}, 5},
		{point{9, 0}, 9},
		{point{0, 1}, 10},
		{point{5, 1}, 15},
		{point{9, 1}, 19},
	}

	for _, tst := range table {
		got := g.xYToI(tst.in)
		if diff := cmp.Diff(tst.want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestContainsSymbol(t *testing.T) {
	table := []struct {
		in   []byte
		want bool
	}{
		{[]byte{'.', '.', '.', '.'}, false},
		{[]byte{'*', '.', '.', '.'}, true},
		{[]byte{'*', '/', '.', '.'}, true},
		{[]byte{'.', '2', '.', '.'}, false},
		{[]byte{'*', '2', '.', '.'}, true},
	}

	for _, tst := range table {
		got := containsSymbol(tst.in)
		if diff := cmp.Diff(tst.want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
