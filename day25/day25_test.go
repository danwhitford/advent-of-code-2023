package day25

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d25 := &Day25{}
	d25.SetInput(exampleIn)
	got := d25.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart25(t *testing.T) {
	exampleIn := ""

	want := ""

	d25 := &Day25{}
	d25.SetInput(exampleIn)
	got := d25.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
