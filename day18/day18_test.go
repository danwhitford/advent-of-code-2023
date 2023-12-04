package day18

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d18 := &Day18{}
	d18.SetInput(exampleIn)
	got := d18.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart18(t *testing.T) {
	exampleIn := ""

	want := ""

	d18 := &Day18{}
	d18.SetInput(exampleIn)
	got := d18.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
