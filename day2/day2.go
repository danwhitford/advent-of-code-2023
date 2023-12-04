package day2

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
)

type Day2 struct {
	games []game
}

type cube struct {
	quantity int
	colour   string
}

type set struct {
	cubes []cube
}

type game struct {
	id   int
	sets []set
}

type parser struct {
	tokens   []helpers.Token
	location int
}

func (p *parser) parseGame() game {
	g := game{}
	p.readAssertLexeme("Game")
	gameIndexToken := p.readAssertType("int")
	gameIndex, _ := strconv.Atoi(gameIndexToken.Val)
	p.readAssertType("colon")

	for p.location < len(p.tokens) {
		g.sets = append(g.sets, p.parseSet())
	}

	g.id = gameIndex
	return g
}

func (p *parser) parseSet() set {
	cc := make([]cube, 0)
	c := cube{}
	for p.location < len(p.tokens) {
		t := p.read()
		switch t.T {
		case "int":
			q, _ := strconv.Atoi(t.Val)
			c.quantity = q
		case "keyword":
			c.colour = t.Val
		case "comma":
			cc = append(cc, c)
			continue
		case "semicolon":
			cc = append(cc, c)
			return set{cc}
		}
	}
	cc = append(cc, c)
	return set{cc}
}

func (p *parser) read() helpers.Token {
	t := p.tokens[p.location]
	p.location++
	return t
}

func (p *parser) readAssertType(expected string) helpers.Token {
	t := p.tokens[p.location]
	if t.T != expected {
		panic(fmt.Errorf("expected %s but got %s", expected, t.T))
	}
	p.location++
	return t
}

func (p *parser) readAssertLexeme(expected string) helpers.Token {
	t := p.tokens[p.location]
	if t.Val != expected {
		panic(fmt.Errorf("expected %s but got %s", expected, t.Val))
	}
	p.location++
	return t
}

func parseLine(s string) game {
	tser := helpers.NewTokeniser(s)
	tokens := tser.Tokenise()
	pser := parser{tokens, 0}
	g := pser.parseGame()
	return g
}

func (day *Day2) SetInput(s string) {
	lines := helpers.GetLines(s)
	for _, line := range lines {
		g := parseLine(line)
		day.games = append(day.games, g)
	}
}

func (c cube) valid() bool {
	switch c.colour {
	case "red":
		return c.quantity <= 12
	case "green":
		return c.quantity <= 13
	case "blue":
		return c.quantity <= 14
	}
	panic(fmt.Errorf("unexpected colour '%+v'", c))
}

func (s set) valid() bool {
	return helpers.Every[cube](s.cubes, func(t cube) bool {
		return t.valid()
	})
}

func (g game) valid() bool {
	return helpers.Every[set](g.sets, func(t set) bool {
		return t.valid()
	})
}

func (day Day2) SolvePart1() string {
	total := 0
	for _, g := range day.games {
		if g.valid() {
			total += g.id
		}
	}
	return fmt.Sprint(total)
}

func (g game) getPowers() int {
	maxes := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, s := range g.sets {
		for _, c := range s.cubes {
			if c.quantity > maxes[c.colour] {
				maxes[c.colour] = c.quantity
			}
		}
	}

	return maxes["red"] * maxes["green"] * maxes["blue"]
}

func (day Day2) SolvePart2() string {
	total := 0
	for _, g := range day.games {
		total += g.getPowers()
	}
	return fmt.Sprint(total)
}
