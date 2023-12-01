package main

import (
	"aoc2023/day1"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"aoc2023/day13"
	"aoc2023/day14"
	"aoc2023/day15"
	"aoc2023/day16"
	"aoc2023/day17"
	"aoc2023/day18"
	"aoc2023/day19"
	"aoc2023/day2"
	"aoc2023/day20"
	"aoc2023/day21"
	"aoc2023/day22"
	"aoc2023/day23"
	"aoc2023/day24"
	"aoc2023/day25"
	"aoc2023/day3"
	"aoc2023/day4"
	"aoc2023/day5"
	"aoc2023/day6"
	"aoc2023/day7"
	"aoc2023/day8"
	"aoc2023/day9"

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
		"day1":  &day1.Day1{},
		"day2":  &day2.Day2{},
		"day3":  &day3.Day3{},
		"day4":  &day4.Day4{},
		"day5":  &day5.Day5{},
		"day6":  &day6.Day6{},
		"day7":  &day7.Day7{},
		"day8":  &day8.Day8{},
		"day9":  &day9.Day9{},
		"day10": &day10.Day10{},
		"day11": &day11.Day11{},
		"day12": &day12.Day12{},
		"day13": &day13.Day13{},
		"day14": &day14.Day14{},
		"day15": &day15.Day15{},
		"day16": &day16.Day16{},
		"day17": &day17.Day17{},
		"day18": &day18.Day18{},
		"day19": &day19.Day19{},
		"day20": &day20.Day20{},
		"day21": &day21.Day21{},
		"day22": &day22.Day22{},
		"day23": &day23.Day23{},
		"day24": &day24.Day24{},
		"day25": &day25.Day25{},
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
