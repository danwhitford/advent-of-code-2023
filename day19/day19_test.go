package day19

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d19 := &Day19{}
	d19.SetInput(exampleIn)
	got := d19.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart19(t *testing.T) {
	exampleIn := ""

	want := ""

	d19 := &Day19{}
	d19.SetInput(exampleIn)
	got := d19.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
