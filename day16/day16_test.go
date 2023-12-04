package day16

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d16 := &Day16{}
	d16.SetInput(exampleIn)
	got := d16.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart16(t *testing.T) {
	exampleIn := ""

	want := ""

	d16 := &Day16{}
	d16.SetInput(exampleIn)
	got := d16.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
