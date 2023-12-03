package day2

import (
	"aoc2023/helpers"
	"fmt"
	"strconv"
	"strings"
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

type tokeniser struct {
	input    string
	location int
}

type token struct {
	t   string
	val string
}

func (t *tokeniser) tokenise() []token {
	retVal := make([]token, 0)
	for t.location < len(t.input) {
		p := t.peek()
		if p == ':' {
			retVal = append(retVal, token{"colon", ":"})
			t.readAssert(':')
		} else if p == ',' {
			retVal = append(retVal, token{"comma", ","})
			t.readAssert(',')
		} else if p == ';' {
			retVal = append(retVal, token{"semicolon", ";"})
			t.readAssert(';')
		} else if p == ' ' {
			t.readAssert(' ')
		} else {
			lexeme := t.readLexeme()
			if _, err := strconv.Atoi(lexeme); err == nil {
				retVal = append(retVal, token{"int", lexeme})
			} else {
				retVal = append(retVal, token{"string", lexeme})
			}
		}
	}
	return retVal
}

func (t *tokeniser) readLexeme() string {
	var s strings.Builder
	for t.location < len(t.input) {
		l := t.peek()
		if l == ':' || l == ',' || l == ';' {
			return s.String()
		}
		t.readAssert(l)
		if l == ' ' || l == '\n' {
			return s.String()
		}
		s.WriteByte(l)
	}
	return s.String()
}

func (t *tokeniser) peek() byte {
	return t.input[t.location]
}

func (t *tokeniser) readAssert(expected byte) byte {
	l := t.input[t.location]
	if l != expected {
		panic(fmt.Errorf("wanted %c but got %c", expected, l))
	}
	t.location++
	return l
}

type parser struct {
	tokens   []token
	location int
}

func (p *parser) parseGame() game {
	g := game{}
	p.readAssertLexeme("Game")
	gameIndexToken := p.readAssertType("int")
	gameIndex, _ := strconv.Atoi(gameIndexToken.val)
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
		switch t.t {
		case "int":
			q, _ := strconv.Atoi(t.val)
			c.quantity = q
		case "string":
			c.colour = t.val
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

func (p *parser) read() token {
	t := p.tokens[p.location]
	p.location++
	return t
}

func (p *parser) readAssertType(expected string) token {
	t := p.tokens[p.location]
	if t.t != expected {
		panic(fmt.Errorf("expected %s but got %s", expected, t.t))
	}
	p.location++
	return t
}

func (p *parser) readAssertLexeme(expected string) token {
	t := p.tokens[p.location]
	if t.val != expected {
		panic(fmt.Errorf("expected %s but got %s", expected, t.val))
	}
	p.location++
	return t
}

func parseLine(s string) game {
	tser := tokeniser{s, 0}
	tokens := tser.tokenise()
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
