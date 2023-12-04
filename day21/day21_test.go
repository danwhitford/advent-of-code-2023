package day21

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d21 := &Day21{}
	d21.SetInput(exampleIn)
	got := d21.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart21(t *testing.T) {
	exampleIn := ""

	want := ""

	d21 := &Day21{}
	d21.SetInput(exampleIn)
	got := d21.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
