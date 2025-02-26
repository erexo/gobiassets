package out

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/erexo/gobiassets/in"
)

type Monster struct {
	Id                 uint16
	BossClass          BossClass
	Name               string
	Level              uint32
	Health             int32
	Experience         uint64
	Speed              int32
	LookType           uint16
	LookHead           uint8
	LookPrimary        uint8
	LookSecondary      uint8
	LookDetails        uint8
	LookAddon          uint8
	AverageDPS         float64
	AverageHPS         float64
	AverageLoot        float64
	AverageLootPer1khp float64
	ExpHpRatio         float64
	Loot               []LootItem
}

type LootItem struct {
	Id       uint16
	Chance   float64
	MinCount uint8
	MaxCount uint8
}

func MonsterHeader() string {
	return `type Monster struct {
	Id                 uint16
	BossClass          BossClass
	Name               string
	Level              uint32
	Health             int32
	Experience         uint64
	Speed              int32
	LookType           uint16
	LookHead           uint8
	LookPrimary        uint8
	LookSecondary      uint8
	LookDetails        uint8
	LookAddon          uint8
	AverageDPS         float64
	AverageHPS         float64
	AverageLoot        float64
	AverageLootPer1khp float64
	ExpHpRatio         float64
	Loot               []LootItem
}

type LootItem struct {
	Id       uint16
	Chance   float64
	MinCount uint8
	MaxCount uint8
}`
}

func GetMonster(id uint16, m *in.Monster, it map[uint16]*Item) *Monster {
	var level uint32
	if flag, ok := m.Flags["level"]; ok {
		lvl, _ := strconv.Atoi(flag)
		level = uint32(lvl)
	}
	bossClass := BossClassNone
	if flag, ok := m.Flags["boss"]; ok {
		switch flag {
		case "regular":
			bossClass = BossClassRegular
		case "daily":
			bossClass = BossClassDaily
		case "mini":
			bossClass = BossClassMini
		}
	}
	lookType, err := strconv.Atoi(strings.Split(m.Look.Type, ";")[0])
	if err != nil {
		panic(err)
	}
	var worth float64
	var items []LootItem
	for _, item := range m.Loot {
		ids := strings.Split(item.Id, ";")
		chance := item.Chance / float64(len(ids))
		for _, idstr := range ids {
			id, err := strconv.Atoi(idstr)
			if err != nil {
				panic(err)
			}
			it, ok := it[uint16(id)]
			if !ok {
				panic(fmt.Sprintf("Unknown item '%d' in monster \"%s\"", id, m.Name))
			}
			if item.MinCount > math.MaxUint8 {
				panic("MinCount")
			}
			if item.MaxCount > math.MaxUint8 {
				panic("MaxCount")
			}
			items = append(items, LootItem{
				Id:       it.ClientId,
				Chance:   chance * 100,
				MinCount: uint8(item.MinCount),
				MaxCount: uint8(item.MaxCount),
			})
			if it.Worth > 0 {
				avgCount := (float64(item.MinCount) + float64(item.MaxCount)) / 2
				worth += chance * float64(it.Worth) * avgCount
			}
		}
	}

	var dps float64
	for _, stage := range m.Stages {
		if stageDps := calculateDmg(stage.Attacks); dps < stageDps {
			dps = stageDps
		}
	}
	dps += calculateDmg(m.Attacks)
	var hps float64
	for _, def := range m.Defenses {
		chs := def.Chance / (float64(def.Interval) / 1000)
		heal := (float64(def.Min) + float64(def.Max)) / 2
		hps += heal * chs
	}

	return &Monster{
		Id:                 id,
		BossClass:          bossClass,
		Name:               Title(strings.TrimSpace(m.Name)),
		Level:              level,
		Health:             int32(m.Health.Now),
		Experience:         uint64(m.Experience),
		Speed:              int32(m.Speed),
		LookType:           uint16(lookType),
		LookHead:           m.Look.Head,
		LookPrimary:        m.Look.Body,
		LookSecondary:      m.Look.Legs,
		LookDetails:        m.Look.Feet,
		LookAddon:          m.Look.Addons,
		AverageDPS:         dps,
		AverageHPS:         hps,
		AverageLoot:        worth,
		AverageLootPer1khp: worth / float64(m.Health.Now) * 1000,
		ExpHpRatio:         float64(m.Experience) / float64(m.Health.Now),
		Loot:               items,
	}
}

func (m *Monster) String() string {
	var items strings.Builder
	if len(m.Loot) > 0 {
		items.WriteString(`[]LootItem{`)
		for i, item := range m.Loot {
			items.WriteString(fmt.Sprintf("{%d, %.3g, %d, %d}", item.Id, item.Chance, item.MinCount, item.MaxCount))
			if i < len(m.Loot)-1 {
				items.WriteString(", ")
			}
		}
		items.WriteByte('}')
	} else {
		items.WriteString("nil")
	}

	return fmt.Sprintf(`&Monster{%d, %d, "%s", %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %.1f, %.1f, %.1f, %.2f, %.3f, %s}`, m.Id, m.BossClass, m.Name, m.Level, m.Health, m.Experience, m.Speed, m.LookType, m.LookHead, m.LookPrimary, m.LookSecondary, m.LookDetails, m.LookAddon, m.AverageDPS, m.AverageHPS, m.AverageLoot, m.AverageLootPer1khp, m.ExpHpRatio, items.String())
}

func calculateDmg(attacks []in.Attack) float64 {
	var dps float64
	for _, atk := range attacks {
		chs := atk.Chance / (float64(atk.Interval) / 1000)
		var dmg float64
		if atk.Skill > 0 && atk.Attack > 0 {
			dmg = (float64(atk.Skill)*(float64(atk.Attack)*0.05) + (float64(atk.Attack) * 0.5)) / 2
			dmg *= 0.9 // ~ -10% for shield/armor block
		} else {
			dmg = -(float64(atk.Min) + float64(atk.Max)) / 2
		}
		dps += dmg * chs
	}
	return dps
}
