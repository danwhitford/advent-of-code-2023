package day7

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d7 := &Day7{}
	d7.SetInput(exampleIn)
	got := d7.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart7(t *testing.T) {
	exampleIn := ""

	want := ""

	d7 := &Day7{}
	d7.SetInput(exampleIn)
	got := d7.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
