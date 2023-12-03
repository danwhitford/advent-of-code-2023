package helpers

import "strings"

func GetLines(in string) []string {
	return strings.Split(in, "\n")
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
