package day11

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d11 := &Day11{}
	d11.SetInput(exampleIn)
	got := d11.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart11(t *testing.T) {
	exampleIn := ""

	want := ""

	d11 := &Day11{}
	d11.SetInput(exampleIn)
	got := d11.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
