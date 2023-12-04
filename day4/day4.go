package day4

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
)

type Day4 struct {
	cards []card
}

type card struct {
	id            int
	intersections int
	quantityHeld  int
}

type set map[int]interface{}

func parseLine(s string) card {
	tker := helpers.NewTokeniser(s)
	tokens := helpers.NewStream[helpers.Token](tker.Tokenise())

	c := card{
		0,
		0,
		1,
	}
	winners := make(set)
	mine := make(set)

	for tokens.HasNext() {
		t := tokens.Next()
		if t.T == "colon" {
			break
		}
	}

	for tokens.HasNext() {
		t := tokens.Next()
		if t.T == "pipe" {
			break
		} else {
			d, err := strconv.Atoi(t.Val)
			if err != nil {
				panic(err)
			}
			winners[d] = struct{}{}
		}
	}

	for tokens.HasNext() {
		t := tokens.Next()
		d, err := strconv.Atoi(t.Val)
		if err != nil {
			panic(err)
		}
		mine[d] = struct{}{}
	}

	c.intersections = getIntersections(winners, mine)
	return c
}

func getIntersections(s1, s2 set) int {
	intersects := 0
	for w := range s1 {
		if _, prs := s2[w]; prs {
			intersects++
		}
	}
	return intersects
}

func (day *Day4) SetInput(s string) {
	lines := helpers.GetLines(s)
	for i, line := range lines {
		card := parseLine(line)
		card.id = i + 1
		day.cards = append(day.cards, card)
	}
}

func (day Day4) SolvePart1() string {
	total := 0
	for _, card := range day.cards {
		intersects := card.intersections
		if intersects > 0 {
			points := 1 << (intersects - 1)
			total += points
		}
	}
	return fmt.Sprint(total)
}

func (day Day4) SolvePart2() string {
	for i := 0; i < len(day.cards); i++ {
		numberIntersects := day.cards[i].intersections
		for j := i + 1; j <= i+numberIntersects; j++ {
			day.cards[j].quantityHeld += day.cards[i].quantityHeld
		}
	}

	total := 0
	for _, card := range day.cards {
		total += card.quantityHeld
	}

	return fmt.Sprint(total)
}
