package main

import (
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

func main() {
    tmpl, err := template.New("file").Parse(fileTemplate)
    if err != nil {
        panic(err)
    }

    for i := 3; i <= 25; i++ {
        fileName := "day" + strconv.Itoa(i) + "/day" + strconv.Itoa(i) + ".go"
        file, err := os.Create(fileName)
        if err != nil {
            panic(err)
        }
        defer file.Close()

        err = tmpl.Execute(file, struct{ Number int }{Number: i})
        if err != nil {
            panic(err)
        }
    }
}
