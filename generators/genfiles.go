package main

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
)

// Template for generating the files
const fileTemplate = `package day{{.Number}}

type Day{{.Number}} struct {}

func (day *Day{{.Number}}) SetInput(s string) {}

func (day Day{{.Number}}) SolvePart1() string {
	return ""
}

func (day Day{{.Number}}) SolvePart2() string {
	return ""
}
`

const testTemplate = `package day{{.Number}}

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolvePart1(t *testing.T) {
	exampleIn := ""

	want := "-"

	d{{.Number}} := &Day{{.Number}}{}
	d{{.Number}}.SetInput(exampleIn)
	got := d{{.Number}}.SolvePart1()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}

func TestSolvePart{{.Number}}(t *testing.T) {
	exampleIn := ""

	want := ""

	d{{.Number}} := &Day{{.Number}}{}
	d{{.Number}}.SetInput(exampleIn)
	got := d{{.Number}}.SolvePart2()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestError() mismatch (-want +got):\n%s", diff)
	}
}
`

func writeTemplateToFile(filename string, i int, tmpl *template.Template) {
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("Skipping '%s' cos it exists\n", filename)
		return
	}
	if !os.IsNotExist(err) {
		panic(err)
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, struct{ Number int }{Number: i})
	if err != nil {
		panic(err)
	}
}

func main() {
	tmpl, err := template.New("file").Parse(fileTemplate)
	if err != nil {
		panic(err)
	}
	testTmpl, err := template.New("test").Parse(testTemplate)
	if err != nil {
		panic(err)
	}

	for i := 3; i <= 25; i++ {
		fileName := "day" + strconv.Itoa(i) + "/day" + strconv.Itoa(i) + ".go"
		testFileName := "day" + strconv.Itoa(i) + "/day" + strconv.Itoa(i) + "_test.go"
		writeTemplateToFile(fileName, i, tmpl)
		writeTemplateToFile(testFileName, i, testTmpl)
	}
}
