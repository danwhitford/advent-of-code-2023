package main

import (
	"aoc2023/day1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Solver func(string) string
type PuzzleDay interface {
	SetInput(string)
	SolvePart1() string
	SolvePart2() string
}

func main() {
	day := os.Args[1]

	register := map[string]PuzzleDay{
		"day1": &day1.Day1{},
	}

	d, prs := register[day]
	if !prs {
		log.Fatalf("Don't know day for %s", day)
	}

	f, err := os.Open(fmt.Sprintf("inputs/%s", day))
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	s := string(bytes)

	d.SetInput(s)
	fmt.Printf("=== RUNNING %s PART 1 ===\n", strings.ToUpper(day))
	res := d.SolvePart1()
	fmt.Printf("=== RUNNING %s PART 2 ===\n", strings.ToUpper(day))
	res2 := d.SolvePart2()

	fmt.Printf("Result for %s part 1: %s\n", day, res)
	fmt.Printf("Result for %s part 2: %s\n", day, res2)
}
