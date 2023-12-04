package helpers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func GetLines(in string) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(in))
	for scanner.Scan() {
		next := scanner.Text()
		lines = append(lines, next)
	}
	return lines
}

func Every[T any](li []T, fn func(t T) bool) bool {
	for _, el := range li {
		if !fn(el) {
			return false
		}
	}
	return true
}

func Some[T any](li []T, fn func(t T) bool) bool {
	for _, el := range li {
		if fn(el) {
			return true
		}
	}
	return false
}

type Tokeniser struct {
	input    string
	location int
}

func NewTokeniser(input string) Tokeniser {
	return Tokeniser{input, 0}
}

type Token struct {
	T   string
	Val string
}

func (t *Tokeniser) Tokenise() []Token {
	retVal := make([]Token, 0)
	for t.location < len(t.input) {
		p := t.peek()
		if p == ':' {
			retVal = append(retVal, Token{"colon", ":"})
			t.readAssert(':')
		} else if p == ',' {
			retVal = append(retVal, Token{"comma", ","})
			t.readAssert(',')
		} else if p == ';' {
			retVal = append(retVal, Token{"semicolon", ";"})
			t.readAssert(';')
		} else if p == '|' {
			retVal = append(retVal, Token{"pipe", "|"})
			t.readAssert('|')
		} else if p == ' ' {
			t.readAssert(' ')
		} else {
			lexeme := t.readLexeme()
			if _, err := strconv.Atoi(lexeme); err == nil {
				retVal = append(retVal, Token{"int", lexeme})
			} else {
				retVal = append(retVal, Token{"keyword", lexeme})
			}
		}
	}
	return retVal
}

func (t *Tokeniser) readLexeme() string {
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

func (t *Tokeniser) peek() byte {
	return t.input[t.location]
}

func (t *Tokeniser) readAssert(expected byte) byte {
	l := t.input[t.location]
	if l != expected {
		panic(fmt.Errorf("wanted %c but got %c", expected, l))
	}
	t.location++
	return l
}

type Stream[T any] struct {
	data []T
	location int
}

func NewStream[T any](t []T) Stream[T] {
	return Stream[T]{t, 0}
}

func (st Stream[T]) HasNext() bool {
	return st.location < len(st.data)
}

func (st *Stream[T]) Next() T {
	t := st.data[st.location]
	st.location++
	return t
}
