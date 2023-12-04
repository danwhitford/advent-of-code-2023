package day12

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d12 := &Day12{}
	d12.SetInput(exampleIn)
	got := d12.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart12(t *testing.T) {
	exampleIn := ""

	want := ""

	d12 := &Day12{}
	d12.SetInput(exampleIn)
	got := d12.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
