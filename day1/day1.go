package day1

import (
	"aoc2023/helpers"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"unicode"
)

type Day1 struct {
	lines []string
}

func (day *Day1) SetInput(input string) {
	day.lines = helpers.GetLines(input)
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

func (day Day1) SolvePart1() string {
	total := 0
	for _, line := range day.lines {
		total += lineValue(line)
	}

	return fmt.Sprintf("%d", total)
}

func stringToDigit(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}
	return s
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
		"[0-9]",
	}
	reMatches := make([]regexp.Regexp, 0)
	for _, m := range matches {
		reMatches = append(reMatches, *regexp.MustCompile(m))
	}

	indexOfMatches := []struct {
		from        regexp.Regexp
		actualMatch string
		loc         []int
	}{}
	for _, re := range reMatches {
		res := re.FindAllStringIndex(line, -1)
		resStrings := re.FindAllString(line, -1)
		for i, loc := range res {
			indexOfMatches = append(indexOfMatches, struct {
				from        regexp.Regexp
				actualMatch string
				loc         []int
			}{
				re,
				resStrings[i],
				loc,
			})
		}
	}

	sort.Slice(indexOfMatches, func(i, j int) bool {
		return indexOfMatches[i].loc[0] < indexOfMatches[j].loc[0]
	})

	d1 := stringToDigit(indexOfMatches[0].actualMatch)
	d2 := stringToDigit(indexOfMatches[len(indexOfMatches)-1].actualMatch)

	d, err := strconv.Atoi(fmt.Sprintf("%s%s", d1, d2))
	if err != nil {
		log.Fatal(err)
	}
	return d
}

func (day Day1) SolvePart2() string {
	total := 0
	for _, line := range day.lines {
		total += lineValue2(line)
	}
	return fmt.Sprintf("%d", total)
}
