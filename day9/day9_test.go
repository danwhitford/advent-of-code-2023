package day9

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d9 := &Day9{}
	d9.SetInput(exampleIn)
	got := d9.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart9(t *testing.T) {
	exampleIn := ""

	want := ""

	d9 := &Day9{}
	d9.SetInput(exampleIn)
	got := d9.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
