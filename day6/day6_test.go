package day6

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d6 := &Day6{}
	d6.SetInput(exampleIn)
	got := d6.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart6(t *testing.T) {
	exampleIn := ""

	want := ""

	d6 := &Day6{}
	d6.SetInput(exampleIn)
	got := d6.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
