package main

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/erexo/gobiaitem/in"
	"github.com/erexo/gobiaitem/out"
)

const (
	dataPath   = `X:\servers\ntsw\data`
	otbFile    = `items\items.otb`
	xmlFile    = `items\items.xml`
	npcDir     = "npc"
	monsterDir = `monster\exp`

	packageName = "assets"

	itemsFileName    = "_items.go"
	monstersFileName = "_monsters.go"
)

func main() {
	fmt.Println("HI")

	prices := readPrices()
	items := saveItems(prices)
	monsters := saveMonsters()

	verifyLoot(items, monsters)
	verifyPrices(items, prices)

	fmt.Println("BYE")
}

func saveItems(prices prices) []*out.Item {
	defer logTime("Items")()

	data, err := ioutil.ReadFile(filepath.Join(dataPath, otbFile))
	if err != nil {
		panic(err)
	}
	serverClient := ReadOtb(data)

	input := readItems()
	itemsToInclude := make(map[uint16]struct{})

	// categories
	var categoriesStr strings.Builder
	for cat := out.ItemCategoryFirst; cat <= out.ItemCategoryLast; cat++ {
		categoriesStr.WriteString(fmt.Sprintf("\tItemCategory(%d): {", cat))
		if it, ok := out.Items[cat]; ok {
			if len(it) > 0 {
				categoriesStr.WriteByte('\n')
			}
			for _, id := range it {
				itemsToInclude[id] = struct{}{}

				categoriesStr.WriteString("\t\titem")
				categoriesStr.WriteString(strconv.Itoa(int(id)))
				categoriesStr.WriteString(",\n")
			}
			if len(it) > 0 {
				categoriesStr.WriteByte('\t')
			}
		}
		categoriesStr.WriteString("},\n")
	}

	items := make([]*out.Item, 0, len(input))
	for _, item := range input {
		if _, ok := itemsToInclude[uint16(item.Id)]; !ok {
			continue
		}
		if client, ok := serverClient[uint16(item.Id)]; ok {
			it := out.NewItem(client, item)
			if it.Worth == 0 {
				if price, ok := prices[item.Id]; ok {
					it.Worth = int64(price)
				}
			}
			items = append(items, it)
		}
	}

	f, err := os.OpenFile(itemsFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	idChars := 1
	if len(items) > 0 {
		idChars += len(strconv.Itoa(int(items[len(items)-1].ServerId)))
	}

	var varStr strings.Builder
	var switchStr strings.Builder
	for _, item := range items {
		// var
		varName := fmt.Sprintf("item%d", item.ServerId)
		varStr.WriteByte('\t')
		varStr.WriteString(varName)
		nameDiff := idChars - len(strconv.Itoa(int(item.ServerId)))
		for i := 0; i < nameDiff; i++ {
			varStr.WriteByte(' ')
		}
		varStr.WriteString("= &")
		varStr.WriteString(item.String())
		varStr.WriteByte('\n')

		// switch
		switchStr.WriteString(fmt.Sprintf("\n\tcase %d:\n\t\treturn %s", item.ClientId, varName))
	}

	if _, err := fmt.Fprintf(f, `// Code generated by "gobiaitem" using 'go generate'. DO NOT EDIT.

package %s

%s

%s

var (
%s)

var ItemsCategory = map[ItemCategory][]*Item{
%s}

func GetItem(clientId uint16) *Item {
	switch clientId {%s
	default:
		return nil
	}
}
`, packageName, out.Category(), out.ItemType(), varStr.String(), categoriesStr.String(), switchStr.String()); err != nil {
		panic(err)
	}
	return items
}

func saveMonsters() []*out.Monster {
	defer logTime("Monsters")()

	input := readMonsters()

	f, err := os.OpenFile(monstersFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	monsters := make([]*out.Monster, len(input))
	variables := make(map[string]string)
	var nameChars int
	var variableChars int
	for i, monster := range input {
		m := out.GetMonster(monster)
		monsters[i] = m
		if len(m.Name) > nameChars {
			nameChars = len(m.Name)
		}
		varName := Variable(m.Name)
		variables[m.Name] = varName
		if len(varName) > variableChars {
			variableChars = len(varName)
		}
	}
	sort.Sort(monsterByLevel(monsters))
	nameChars++
	variableChars++

	var varStr strings.Builder
	var byLevelStr strings.Builder
	var switchStr strings.Builder
	for _, monster := range monsters {
		varName := variables[monster.Name]

		// var
		varStr.WriteString(fmt.Sprintf("\t%s", varName))
		nameDiff := variableChars - len(varName)
		for i := 0; i < nameDiff; i++ {
			varStr.WriteByte(' ')
		}
		varStr.WriteString(fmt.Sprintf("= %s\n", monster.String()))

		// level
		byLevelStr.WriteString(fmt.Sprintf("\t%s,\n", varName))

		// switch
		switchStr.WriteString(fmt.Sprintf("\n\tcase \"%s\":\n\t\treturn %s", monster.Name, varName))
	}
	if _, err := fmt.Fprintf(f, `// Code generated by "gobiaitem" using 'go generate'. DO NOT EDIT.

package %s

%s

var (
%s)

var MonstersByLevel = []*Monster{
%s}

func GetMonster(name string) *Monster {
	switch name {%s
	default:
		return nil
	}
}
`, packageName, out.MonsterType(), varStr.String(), byLevelStr.String(), switchStr.String()); err != nil {
		panic(err)
	}
	return monsters
}

func verifyLoot(items []*out.Item, monsters []*out.Monster) {
	desired := make(map[uint16]struct{})
	for _, monster := range monsters {
		for _, loot := range monster.Loot {
			desired[loot.Id] = struct{}{}
		}
	}
	for _, item := range items {
		delete(desired, item.ServerId)
	}
	if len(desired) == 0 {
		return
	}

	var sb strings.Builder
	sb.WriteString("Omitted Loot item ids:\n")
	for id := range desired {
		fmt.Fprintf(&sb, "%d, ", id)
	}
	log.Println(sb.String())
}

func verifyPrices(items []*out.Item, prices prices) {
	desired := make(map[uint16]struct{})
	for id := range prices {
		desired[uint16(id)] = struct{}{}
	}
	for _, item := range items {
		delete(desired, item.ServerId)
	}
	if len(desired) == 0 {
		return
	}

	var sb strings.Builder
	sb.WriteString("Omitted Price item ids:\n")
	for id := range desired {
		fmt.Fprintf(&sb, "%d, ", id)
	}
	log.Println(sb.String())
}

func readItems() []*in.Item {
	f, err := os.Open(filepath.Join(dataPath, xmlFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	items := &in.Items{}
	if err := xml.NewDecoder(f).Decode(items); err != nil {
		panic(err)
	}

	return items.Items
}

type prices map[int]int

func readPrices() prices {
	ret := make(prices)
	if err := filepath.WalkDir(filepath.Join(dataPath, npcDir), func(path string, e fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !e.IsDir() && filepath.Ext(path) == ".xml" {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			var n struct {
				Parameters []struct {
					Key   string `xml:"key,attr"`
					Value string `xml:"value,attr"`
				} `xml:"parameters>parameter"`
			}
			err = xml.NewDecoder(f).Decode(&n)
			if err != nil {
				return err
			}
			f.Close()

			for _, p := range n.Parameters {
				switch p.Key {
				case "shop_buyable", "shop_sellable":
					items := strings.Split(p.Value, ";")
					for _, item := range items {
						if item == "" {
							continue
						}
						values := strings.Split(item, ",")
						if len(values) < 3 {
							panic(fmt.Sprintf("Invalid values parameters for '%s' in %s", item, path))
						}
						//name := values[0]
						id, err := strconv.Atoi(values[1])
						if err != nil {
							panic(err)
						}
						cost, err := strconv.Atoi(values[2])
						if err != nil {
							panic(err)
						}
						if v, ok := ret[id]; !ok || v < cost {
							ret[id] = cost
						}
					}
				}
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}
	return ret
}

func readMonsters() []*in.Monster {
	var monsters []*in.Monster
	if err := filepath.WalkDir(filepath.Join(dataPath, monsterDir), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || filepath.Ext(path) != ".xml" {
			return nil
		}
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		m := in.NewMonster()
		if err := xml.NewDecoder(f).Decode(m); err != nil {
			return err
		}
		monsters = append(monsters, m)
		return nil
	}); err != nil {
		panic(err)
	}
	return monsters
}

type monsterByLevel []*out.Monster

func (s monsterByLevel) Len() int      { return len(s) }
func (s monsterByLevel) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s monsterByLevel) Less(i, j int) bool {
	if s[i].Level == s[j].Level {
		return s[i].Health < s[j].Health
	}
	return s[i].Level < s[j].Level
}
