package helpers

import (
	"bufio"
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
