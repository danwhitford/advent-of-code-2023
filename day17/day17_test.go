package day17

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d17 := &Day17{}
	d17.SetInput(exampleIn)
	got := d17.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart17(t *testing.T) {
	exampleIn := ""

	want := ""

	d17 := &Day17{}
	d17.SetInput(exampleIn)
	got := d17.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
