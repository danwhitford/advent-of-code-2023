package day23

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d23 := &Day23{}
	d23.SetInput(exampleIn)
	got := d23.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart23(t *testing.T) {
	exampleIn := ""

	want := ""

	d23 := &Day23{}
	d23.SetInput(exampleIn)
	got := d23.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
