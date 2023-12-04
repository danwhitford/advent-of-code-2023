package day8

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d8 := &Day8{}
	d8.SetInput(exampleIn)
	got := d8.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart8(t *testing.T) {
	exampleIn := ""

	want := ""

	d8 := &Day8{}
	d8.SetInput(exampleIn)
	got := d8.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
