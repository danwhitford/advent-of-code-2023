package day15

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d15 := &Day15{}
	d15.SetInput(exampleIn)
	got := d15.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart15(t *testing.T) {
	exampleIn := ""

	want := ""

	d15 := &Day15{}
	d15.SetInput(exampleIn)
	got := d15.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
