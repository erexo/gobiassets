package out

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erexo/gobiaitem/in"
)

type Monster struct {
	Name       string
	Level      uint32
	Health     int32
	Experience uint64
	Speed      int32
	LookType   uint32
	AverageDPS float64
	AverageHPS float64
	Loot       []LootItem
}

type LootItem struct {
	Id       uint16
	Chance   float64
	MinCount uint16
	MaxCount uint16
}

func MonsterType() string {
	return `type Monster struct {
	Name       string
	Level      uint32
	Health     int32
	Experience uint64
	Speed      int32
	LookType   uint32
	AverageDPS float64
	AverageHPS float64
	Loot       []Item
}

type Item struct {
	Id       uint16
	Chance   float64
	MinCount uint16
	MaxCount uint16
}`
}

func GetMonster(m *in.Monster) *Monster {
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
		items = append(items, LootItem{
			Id:       uint16(item.Id),
			Chance:   item.Chance,
			MinCount: uint16(item.MinCount),
			MaxCount: uint16(item.MaxCount),
		})
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
		Name:       Title(m.Name),
		Level:      level,
		Health:     int32(m.Health.Now),
		Experience: uint64(m.Experience),
		Speed:      int32(m.Speed),
		LookType:   uint32(lookType),
		AverageDPS: dps,
		AverageHPS: hps,
		Loot:       items,
	}
}

func (m *Monster) String() string {
	var items strings.Builder
	items.WriteString(`[]Item{`)
	for i, item := range m.Loot {
		items.WriteString(fmt.Sprintf("{%d, %.5f, %d, %d}", item.Id, item.Chance, item.MinCount, item.MaxCount))
		if i < len(m.Loot)-1 {
			items.WriteString(", ")
		}
	}
	items.WriteByte('}')

	return fmt.Sprintf(`&Monster{"%s", %d, %d, %d, %d, %d, %.2f, %.2f, %s}`, m.Name, m.Level, m.Health, m.Experience, m.Speed, m.LookType, m.AverageDPS, m.AverageHPS, items.String())
}
