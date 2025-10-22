package in

import (
	"bufio"
	"encoding/xml"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Spells struct {
	Instant []*Spell `xml:"instant"`
	Conjure []*Spell `xml:"conjure"`

	Passive []*Spell `xml:"passive"`
}

type Spell struct {
	Words                   string          `xml:"words,attr"`
	Level                   int             `xml:"level,attr"`
	Chakra                  int             `xml:"chakra,attr"`
	ChakraPercent           int             `xml:"chakrapercent,attr"`
	Health                  int             `xml:"health,attr"`
	Soul                    int             `xml:"soul,attr"`
	IsSpecial               BoolInt         `xml:"isspecial,attr"`
	Range                   int             `xml:"range,attr"`
	Direction               BoolInt         `xml:"direction,attr"`
	Param                   BoolInt         `xml:"param,attr"`
	NeedTarget              BoolInt         `xml:"needTarget,attr"`
	IsBuff                  BoolInt         `xml:"isbuff,attr"`
	CasterTargetOrDirection BoolInt         `xml:"casterTargetOrDirection,attr"`
	ConjureId               int             `xml:"conjureId,attr"`
	SelfTarget              BoolInt         `xml:"selftarget,attr"`
	Enabled                 string          `xml:"enabled,attr"`
	Aggressive              string          `xml:"aggressive,attr"`
	Exhaustion              int             `xml:"exhaustion,attr"`
	SType                   int             `xml:"sType,attr"`
	SExhaustion             int             `xml:"sExhaustion,attr"`
	Script                  string          `xml:"script,attr"`
	Vocations               []SpellVocation `xml:"vocation"`

	Hide        BoolInt `xml:"_hide,attr"`
	Type        string  `xml:"_type,attr"`
	Description string  `xml:"_desc,attr"`

	Attack float64 `xml:"_attk,attr"`
	Heal   float64
	Reduce float64
}

type SpellVocation struct {
	Profession  string `xml:"profession,attr"`
	Role        string `xml:"role,attr"`
	Description string `xml:"_desc,attr"`
}

type BoolInt bool

func (b *BoolInt) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "1" {
		*b = true
	} else {
		*b = false
	}
	return nil
}

func ReadSpellsFile() map[string][]*Spell {
	f, err := os.Open(filepath.Join(dataPath, spellsXmlFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var s Spells
	if err := xml.NewDecoder(f).Decode(&s); err != nil {
		panic(err)
	}

	spells := append(s.Instant, s.Conjure...)
	spells = append(spells, s.Passive...)

	// remove disabled spells
	filtered := spells[:0]
	for _, s := range spells {
		if s.Enabled == "0" || s.Hide {
			continue
		}
		filtered = append(filtered, s)
	}
	spells = filtered

	ret := make(map[string][]*Spell)
	roleSpells := make(map[string][]*Spell)

	lib := ReadSpellsLibFile()
	for _, spell := range spells {
		readSpellFile(spell, lib)

		const ALL = "All"
		if len(spell.Vocations) == 0 {
			ret[ALL] = append(ret[ALL], spell)
		}

		for _, voc := range spell.Vocations {
			if voc.Profession != "" {
				ret[voc.Profession] = append(ret[voc.Profession], spell)
			}
			if voc.Role != "" {
				roleSpells[voc.Role] = append(roleSpells[voc.Role], spell)
			}
		}
	}

	vocs := ReadVocationsFile()
	for name := range ret {
		for _, voc := range vocs {
			if name == voc.Name {
				if roleSpells, ok := roleSpells[voc.Formula.MainRole]; ok {
					ret[name] = append(ret[name], roleSpells...)
				}
				break
			}
		}
		sort.Slice(ret[name], func(i, j int) bool {
			return ret[name][i].Level < ret[name][j].Level
		})
	}

	return ret
}

func ReadSpellsLibFile() map[string]float64 {
	f, err := os.Open(filepath.Join(dataPath, spellsLibFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ret := make(map[string]float64)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip empty lines or separators
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-") {
			continue
		}

		// split at '='
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // skip malformed lines
		}

		key := strings.TrimSpace(parts[0])
		valueStr := strings.TrimSpace(parts[1])

		if key == "" || valueStr == "" {
			continue
		}

		// parse float64
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(err)
		}

		ret[key] = value
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return ret
}

func readSpellFile(spell *Spell, lib map[string]float64) {
	if spell.Script == "" {
		return
	}

	data, err := os.ReadFile(filepath.Join(dataPath, spellsDir, spell.Script))
	if err != nil {
		panic(err)
	}
	script := string(data)

	spellValues := make(map[string]float64)

	for k, v := range lib {
		// split prefix (e.g. ATTACK_10_TARGET → ATTACK)
		parts := strings.SplitN(k, "_", 2)
		if len(parts) < 2 {
			continue
		}
		prefix := parts[0]

		if _, ok := spellValues[prefix]; ok {
			continue
		}

		if strings.Contains(script, k) {
			spellValues[prefix] = v
		}
	}

	if atk, ok := spellValues["ATTACK"]; ok {
		spell.Attack = atk
	}
	if atk, ok := spellValues["HEAL"]; ok {
		spell.Heal = atk
	}
	if atk, ok := spellValues["REDUCE"]; ok {
		spell.Reduce = atk
	}
}
