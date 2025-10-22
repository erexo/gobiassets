package out

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"

	"github.com/erexo/gobiassets/in"
)

func SaveSpells() {
	spells := in.ReadSpellsFile()

	_ = os.MkdirAll(outputDir, 0755)
	f, err := os.OpenFile(path.Join(outputDir, professionsFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	keys := make([]string, 0, len(spells))
	for k := range spells {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		for _, spell := range spells[k] {
			chakra := strconv.Itoa(spell.Chakra)
			if spell.ChakraPercent > 0 {
				chakra = fmt.Sprintf("%d%%", spell.ChakraPercent)
			}
			if spell.Health > 0 {
				chakra = fmt.Sprintf("%d health", spell.Health)
			}

			attack := fmt.Sprintf("%.f", spell.Attack)
			if spell.Heal > 0 {
				attack = fmt.Sprintf("%.f%%", spell.Heal)
			}

			reduce := strings.TrimSuffix(fmt.Sprintf("%.1f", spell.Reduce), ".0")
			exhaust := strings.TrimSuffix(fmt.Sprintf("%.1f", getSpellExhaust(spell)), ".0")

			description := spell.Description
			for _, voc := range spell.Vocations {
				if voc.Profession == k {
					if voc.Description != "" {
						description = voc.Description
					}
					break
				}
			}

			fmt.Fprintf(f, `array("%s", "%s", "%s", "%d", "%s", "%d", "%s", "%s", "%s", "%s"),
`, k, spellName(spell.Words), getSpellType(spell), spell.Level, chakra, spell.Soul, attack, reduce, exhaust, description)
		}
	}
}

func spellName(s string) string {
	words := strings.Fields(s) // split string by spaces
	for i, word := range words {
		if strings.ToLower(word) != "no" {
			words[i] = strings.Title(word) // capitalize first letter
		} else {
			words[i] = "no" // keep "no" lowercase
		}
	}
	return strings.Join(words, " ")
}

func getSpellExhaust(s *in.Spell) float64 {
	if s.Type == "Passive" {
		return 0
	}

	if s.SExhaustion > 0 {
		return float64(s.SExhaustion) / 1000
	}
	if s.Exhaustion > 0 {
		return float64(s.Exhaustion) / 1000
	}

	if s.Aggressive != "0" {
		switch s.Level {
		case 10:
			return 4.
		case 20:
			return 4.
		case 40:
			return 6.
		case 70:
			return 6.
		case 90, 100:
			return 6.
		case 110:
			return 6.
		case 160:
			return 6.
		case 220:
			return 6.
		case 330:
			return 6.
		}
	}

	return 1.
}

func getSpellType(s *in.Spell) string {
	if s.Type != "" {
		return s.Type
	}

	if s.ConjureId > 0 {
		return "Conjure"
	}

	if s.Heal > 0 {
		return "Healing"
	}

	if s.Direction {
		return "Wave"
	}

	if s.SelfTarget {
		return "Self"
	}

	if s.IsSpecial && s.Level == 50 {
		return "Buff"
	}

	if s.Param && s.IsBuff {
		return "Self/Target"
	}

	if s.Param || s.NeedTarget || s.CasterTargetOrDirection || s.Range > 0 {
		return "Target"
	}
	// Wave, Area, Target, []Support, []Healing, []Clone, Buff, Party, []Conjure, Self
	return "Area"
}
