package day4

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d4 := &Day4{}
	d4.SetInput(exampleIn)
	got := d4.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart4(t *testing.T) {
	exampleIn := ""

	want := ""

	d4 := &Day4{}
	d4.SetInput(exampleIn)
	got := d4.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
