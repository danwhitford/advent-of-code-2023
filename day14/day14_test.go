package day14

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d14 := &Day14{}
	d14.SetInput(exampleIn)
	got := d14.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart14(t *testing.T) {
	exampleIn := ""

	want := ""

	d14 := &Day14{}
	d14.SetInput(exampleIn)
	got := d14.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
