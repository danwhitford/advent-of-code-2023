package day13

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d13 := &Day13{}
	d13.SetInput(exampleIn)
	got := d13.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart13(t *testing.T) {
	exampleIn := ""

	want := ""

	d13 := &Day13{}
	d13.SetInput(exampleIn)
	got := d13.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
