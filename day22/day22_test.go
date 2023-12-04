package day22

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := ""

	d22 := &Day22{}
	d22.SetInput(exampleIn)
	got := d22.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart22(t *testing.T) {
	exampleIn := ""

	want := ""

	d22 := &Day22{}
	d22.SetInput(exampleIn)
	got := d22.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
