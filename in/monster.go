package in

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
)

const (
	maxSpellChance = 100
	maxItemChance  = 50000
)

func ReadMonsterFile() map[string]IdName {
	f, err := os.Open(filepath.Join(dataPath, monsterXmlFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var m struct {
		Monsters []struct {
			Id   uint16 `xml:"id,attr"`
			Name string `xml:"name,attr"`
			File string `xml:"file,attr"`
		} `xml:"monster"`
	}
	if err := xml.NewDecoder(f).Decode(&m); err != nil {
		panic(err)
	}
	ret := make(map[string]IdName)
	for _, monster := range m.Monsters {
		ret[monster.File] = IdName{monster.Id, monster.Name}
	}
	return ret
}

func ReadMonsters(dir string) []*Monster {
	basePath := filepath.Join(dataPath, monsterDir, dir)
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic(err)
	}

	var monsters []*Monster
	for _, file := range files {
		fpath := filepath.Join(basePath, file.Name())
		if file.IsDir() || filepath.Ext(fpath) != ".xml" {
			continue
		}
		f, err := os.Open(fpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		m := NewMonster(path.Join(dir, filepath.Base(fpath)))
		if err := xml.NewDecoder(f).Decode(m); err != nil {
			panic(err)
		}
		monsters = append(monsters, m)
	}
	return monsters
}

type IdName struct {
	Id   uint16
	Name string
}

type Monster struct {
	Path       string
	Name       string   `xml:"name,attr"`
	Experience int      `xml:"experience,attr"`
	Speed      int      `xml:"speed,attr"`
	Health     Health   `xml:"health"`
	Look       Look     `xml:"look"`
	Flags      Flags    `xml:"flags>flag"`
	Attacks    []Attack `xml:"attacks>attack"`
	Stages     []Stage  `xml:"attacks>stage"`
	Defenses   Defenses `xml:"defenses"`
	Loot       Loot     `xml:"loot>item"`
}

func NewMonster(path string) *Monster {
	return &Monster{
		Path:  path,
		Flags: make(Flags),
	}
}

type Health struct {
	Now   int `xml:"now,attr"`
	Armor int `xml:"armor,attr"`
}

type Look struct {
	Type   string `xml:"type,attr"`
	Head   uint8  `xml:"head,attr"`
	Body   uint8  `xml:"body,attr"`
	Legs   uint8  `xml:"legs,attr"`
	Feet   uint8  `xml:"feet,attr"`
	Addons uint8  `xml:"addons,attr"`
}

type Flags map[string]string

func (f Flags) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		f[attr.Name.Local] = attr.Value
	}
	dec.Skip()
	return nil
}

type Stage struct {
	Id      int      `xml:"id,attr"`
	Attacks []Attack `xml:"attack"`
}

type Attack struct {
	Name      string  `xml:"name,attr"`
	Interval  int     `xml:"interval,attr"`
	Chance    float64 `xml:"chance,attr"`
	Skill     int     `xml:"skill,attr"`
	Attack    int     `xml:"attack,attr"`
	Min       int     `xml:"min,attr"`
	Max       int     `xml:"max,attr"`
	Radius    int     `xml:"radius,attr"`
	Target    int     `xml:"target,attr"`
	Range     int     `xml:"range,attr"`
	Length    int     `xml:"length,attr"`
	Spread    int     `xml:"spread,attr"`
	DelayType int     `xml:"delayType,attr"`
}

func (a *Attack) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type atk Attack
	new := atk{
		Chance: maxSpellChance,
	}
	if err := dec.DecodeElement(&new, &start); err != nil {
		return err
	}
	new.Chance = clamp(new.Chance, 0, maxSpellChance) / maxSpellChance
	*a = (Attack)(new)
	return nil
}

type Defenses struct {
	Defenses []Defense `xml:"defense"`
}

type Defense struct {
	Name     string  `xml:"name,attr"`
	Interval int     `xml:"interval,attr"`
	Chance   float64 `xml:"chance,attr"`
	Min      int     `xml:"min,attr"`
	Max      int     `xml:"max,attr"`
}

func (d *Defense) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type def Defense
	new := def{
		Chance: maxSpellChance,
	}
	if err := dec.DecodeElement(&new, &start); err != nil {
		return err
	}
	new.Chance = clamp(new.Chance, 0, maxSpellChance) / maxSpellChance
	*d = (Defense)(new)
	return nil
}

type Loot []LootItem

func (i *Loot) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var s []itemInside
	if err := dec.DecodeElement(&s, &start); err != nil {
		return err
	}
	*i = append(*i, getItems(s, 1)...)
	sort.Sort(*i)
	return nil
}

func (l Loot) Len() int      { return len(l) }
func (l Loot) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l Loot) Less(i, j int) bool {
	if l[i].Chance == l[j].Chance {
		return l[i].Id < l[j].Id
	}
	return l[i].Chance > l[j].Chance
}

func getItems(items []itemInside, chanceMul float64) []LootItem {
	ret := make([]LootItem, 0, len(items))
	for _, item := range items {
		item.Chance *= chanceMul
		ret = append(ret, item.LootItem)
		if len(item.Inside.Items) > 0 {
			ret = append(ret, getItems(item.Inside.Items, item.Chance)...)
		}
	}
	return ret
}

type LootItem struct {
	Id       string  `xml:"id,attr"`
	Chance   float64 `xml:"chance,attr"`
	MinCount int     `xml:"mincount,attr"`
	MaxCount int     `xml:"count,attr"`

	SharedMultiper float64 `xml:"sharedmultiper,attr"`
	ChanceMultiper float64 `xml:"chancemultiper,attr"`
	MaxChance      float64 `xml:"maxchance,attr"`
}

type itemInside struct {
	LootItem
	Inside struct {
		Items []itemInside `xml:"item"`
	} `xml:"inside"`
	Attributes []xml.Attr `xml:",any,attr"`
}

func (i *itemInside) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	type itIns itemInside
	new := itIns{
		LootItem: LootItem{
			Chance:    maxItemChance,
			MaxChance: maxItemChance,
		},
	}
	if err := dec.DecodeElement(&new, &start); err != nil {
		return err
	}
	for _, attr := range new.Attributes {
		switch attr.Name.Local {
		case "chance1":
			if v, err := strconv.Atoi(attr.Value); err == nil {
				new.LootItem.Chance = float64(v)
			}
		case "countmin":
			if v, err := strconv.Atoi(attr.Value); err == nil {
				new.LootItem.MinCount = v
			}
		case "countmax":
			if v, err := strconv.Atoi(attr.Value); err == nil {
				new.LootItem.MaxCount = v
			}
		}
	}
	new.MinCount = clamp(new.MinCount, 1, 100)
	new.MaxCount = clamp(new.MaxCount, 1, 100)
	new.Chance = clamp(new.Chance, 0, maxItemChance) / maxItemChance
	new.MaxChance = clamp(new.MaxChance, 0, maxItemChance) / maxItemChance
	*i = itemInside(new)
	return nil
}
