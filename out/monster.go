package out

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erexo/gobiassets/in"
)

type Monster struct {
	Id            uint16
	Name          string
	Level         uint32
	Health        int32
	Experience    uint64
	Speed         int32
	LookType      uint16
	LookHead      uint8
	LookPrimary   uint8
	LookSecondary uint8
	LookDetails   uint8
	LookAddon     uint8
	AverageDPS    float64
	AverageHPS    float64
	Loot          []LootItem
}

type LootItem struct {
	Id       uint16
	Chance   float64
	MinCount uint16
	MaxCount uint16
}

func MonsterType() string {
	return `type Monster struct {
	Id            uint16
	Name          string
	Level         uint32
	Health        int32
	Experience    uint64
	Speed         int32
	LookType      uint16
	LookHead      uint8
	LookPrimary   uint8
	LookSecondary uint8
	LookDetails   uint8
	LookAddon     uint8
	AverageDPS    float64
	AverageHPS    float64
	Loot          []LootItem
}

type LootItem struct {
	Id       uint16
	Chance   float64
	MinCount uint16
	MaxCount uint16
}`
}

func GetMonster(id uint16, m *in.Monster) *Monster {
	var level uint32
	if lvl, ok := m.Flags["level"]; ok {
		level = uint32(lvl)
	}
	lookType, err := strconv.Atoi(strings.Split(m.Look.Type, ";")[0])
	if err != nil {
		panic(err)
	}
	var items []LootItem
	for _, item := range m.Loot {
		ids := strings.Split(item.Id, ";")
		chance := item.Chance / float64(len(ids))
		for _, idstr := range ids {
			id, err := strconv.Atoi(idstr)
			if err != nil {
				panic(err)
			}
			items = append(items, LootItem{
				Id:       uint16(id),
				Chance:   chance,
				MinCount: uint16(item.MinCount),
				MaxCount: uint16(item.MaxCount),
			})
		}
	}

	var dps, hps float64
	for _, atk := range m.Attacks {
		chs := atk.Chance / (float64(atk.Interval) / 1000)
		var dmg float64
		if atk.Skill > 0 && atk.Attack > 0 {
			dmg = (float64(atk.Skill)*(float64(atk.Attack)*0.05) + (float64(atk.Attack) * 0.5)) / 2
			dmg *= 0.9 // ~ -10% for shield/armor block
		} else {
			dmg = -(float64(atk.Min) + float64(atk.Max)) / 2
		}
		// (int32_t)std::ceil((attackSkill * (attackValue * 0.05)) + (attackValue * 0.5));
		dps += dmg * chs
	}
	for _, def := range m.Defenses {
		chs := def.Chance / (float64(def.Interval) / 1000)
		heal := (float64(def.Min) + float64(def.Max)) / 2
		hps += heal * chs
	}

	return &Monster{
		Id:            id,
		Name:          Title(m.Name),
		Level:         level,
		Health:        int32(m.Health.Now),
		Experience:    uint64(m.Experience),
		Speed:         int32(m.Speed),
		LookType:      uint16(lookType),
		LookHead:      m.Look.Head,
		LookPrimary:   m.Look.Body,
		LookSecondary: m.Look.Legs,
		LookDetails:   m.Look.Feet,
		LookAddon:     m.Look.Addons,
		AverageDPS:    dps,
		AverageHPS:    hps,
		Loot:          items,
	}
}

func (m *Monster) String() string {
	var items strings.Builder
	if len(m.Loot) > 0 {
		items.WriteString(`[]LootItem{`)
		for i, item := range m.Loot {
			items.WriteString(fmt.Sprintf("{%d, %.5f, %d, %d}", item.Id, item.Chance, item.MinCount, item.MaxCount))
			if i < len(m.Loot)-1 {
				items.WriteString(", ")
			}
		}
		items.WriteByte('}')
	} else {
		items.WriteString("nil")
	}

	return fmt.Sprintf(`&Monster{%d, "%s", %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %.2f, %.2f, %s}`, m.Id, m.Name, m.Level, m.Health, m.Experience, m.Speed, m.LookType, m.LookHead, m.LookPrimary, m.LookSecondary, m.LookDetails, m.LookAddon, m.AverageDPS, m.AverageHPS, items.String())
}
