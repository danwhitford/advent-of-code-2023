package day24

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d24 := &Day24{}
	d24.SetInput(exampleIn)
	got := d24.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart24(t *testing.T) {
	exampleIn := ""

	want := ""

	d24 := &Day24{}
	d24.SetInput(exampleIn)
	got := d24.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
