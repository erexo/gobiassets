package out

import (
	"strings"
	"unicode"
)

func Title(s string) string {
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if !unicode.IsLetter(prev) {
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