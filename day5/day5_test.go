package day5

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d5 := &Day5{}
	d5.SetInput(exampleIn)
	got := d5.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart5(t *testing.T) {
	exampleIn := ""

	want := ""

	d5 := &Day5{}
	d5.SetInput(exampleIn)
	got := d5.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
