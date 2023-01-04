package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func logTime(msg string) func() {
	now := time.Now()
	return func() {
		fmt.Printf("%s in %s\n", msg, time.Now().Sub(now))
	}
}

func Variable(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			}
			if first {
				first = false
				return unicode.ToLower(r)
			}
			return r
		},
		s)
}
