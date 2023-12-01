package helpers

import "strings"

func GetLines(in string) []string {
	return strings.Split(in, "\n")
}
