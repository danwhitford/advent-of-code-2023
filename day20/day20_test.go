package day20

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d20 := &Day20{}
	d20.SetInput(exampleIn)
	got := d20.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart20(t *testing.T) {
	exampleIn := ""

	want := ""

	d20 := &Day20{}
	d20.SetInput(exampleIn)
	got := d20.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
