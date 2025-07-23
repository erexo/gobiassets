package out

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func LogTime(msg string) func() {
	now := time.Now()
	return func() {
		fmt.Printf("%s\tin %s\n", msg, time.Now().Sub(now))
	}
}

func Variable(s string) string {
	first := true
	return strings.Map(
		func(r rune) rune {
			if unicode.IsSpace(r) {
				first = true
				return -1
			}
			if first {
				first = false
				return unicode.ToUpper(r)
			}
			return unicode.ToLower(r)
		},
		s)
}

func Title(s string) string {
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if !unicode.IsLetter(prev) && prev != '\'' {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return unicode.ToLower(r)
		},
		s)
}

func Role(s string) ItemRole {
	switch s {
	case "ninjutsu":
		return ItemRoleNinjutsu
	case "weapons":
		return ItemRoleWeapons
	case "defense":
		return ItemRoleDefense
	default:
		return ItemRoleAll
	}
}

func Type(s string) ItemType {
	switch s {
	case "boss":
		return ItemTypeBoss
	case "mission":
		return ItemTypeMission
	default:
		return ItemTypeNone
	}
}
