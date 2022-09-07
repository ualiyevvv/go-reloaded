package goreloaded

import (
	"strings"
)

func Split(text string) []string {
	newtxt := strings.Split(text, "\n")

	return newtxt
}
