package day1

import (
	"aoc2023/helpers"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"
)

type Day1 struct {
	Input string
}

func (day *Day1) SetInput(input string) {
	day.Input = input
}

func lineValue(line string) int {
	var first, last rune

	for _, char := range line {
		if unicode.IsDigit(char) {
			if first == 0 {
				first = char
			}
			last = char
		}
	}

	d, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
	if err != nil {
		panic(err)
	}
	return d
}

func (day Day1) SolvePart1 () string {
	lines := helpers.GetLines(day.Input)
	total := 0
	for _, line := range lines {
		total += lineValue(line)
	}

	return fmt.Sprintf("%d", total)
}

func lineValue2(line string) int {
	matches := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	reMatches := make([]regexp.Regexp, 0)
	for _, m := range matches {
		reMatches = append(reMatches, *regexp.MustCompile(m))
	}

	indexOfMatches := []struct {
		from regexp.Regexp
		to   int
		locs [][]int
	}{}
	for i, re := range reMatches {
		res := re.FindAllStringIndex(line, -1)
		if res != nil {
			indexOfMatches = append(indexOfMatches, struct {
				from regexp.Regexp
				to   int
				locs [][]int
			}{
				re,
				i,
				res,
			})
		}
	}

	var firstMatch, lastMatch *struct {
		from regexp.Regexp
		to   int
		loc  int
	}
	for i, match := range indexOfMatches {
		if i == 0 || match.locs[0][0] < firstMatch.loc {
			firstMatch = &struct {
				from regexp.Regexp
				to   int
				loc  int
			}{
				match.from,
				match.to,
				match.locs[0][0],
			}
		}
		if i == 0 || match.locs[len(match.locs)-1][0] > lastMatch.loc {
			lastMatch = &struct {
				from regexp.Regexp
				to   int
				loc  int
			}{
				match.from,
				match.to,
				match.locs[len(match.locs)-1][0],
			}
		}
	}

	var firstString, lastString = line, line
	if firstMatch != nil {
		firstString = firstMatch.from.ReplaceAllString(line, fmt.Sprintf("%d", firstMatch.to))
	}
	if lastMatch != nil {
		lastString = lastMatch.from.ReplaceAllString(line, fmt.Sprintf("%d", lastMatch.to))
	}

	var first, last rune
	for _, c := range firstString {
		if unicode.IsDigit(c) && first == 0{
			first = c
			break
		}
	}
	for _, c := range lastString {
		if unicode.IsDigit(c) {
			last = c
		}
	}

	d, err := strconv.Atoi(fmt.Sprintf("%c%c", first, last))
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func (day Day1) SolvePart2() string {
	lines := helpers.GetLines(day.Input)
	total := 0
	for _, line := range lines {
		total += lineValue2(line)
	}
	return fmt.Sprintf("%d", total)
}
