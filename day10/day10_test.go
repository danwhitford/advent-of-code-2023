package day10

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d10 := &Day10{}
	d10.SetInput(exampleIn)
	got := d10.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart10(t *testing.T) {
	exampleIn := ""

	want := ""

	d10 := &Day10{}
	d10.SetInput(exampleIn)
	got := d10.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
