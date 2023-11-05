package main

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/erexo/gobiassets/in"
	"github.com/erexo/gobiassets/out"
)

const (
	dataPath       = `X:\servers\ntsw\data`
	itemsOtbFile   = `items\items.otb`
	itemsXmlFile   = `items\items.xml`
	npcDir         = "npc"
	monsterDir     = `monster`
	monsterXmlFile = `monster\monsters.xml`

	packageName = "assets"

	itemsFileName    = "_items.go"
	monstersFileName = "_monsters.go"
)

func main() {
	fmt.Println("HI")

	prices := readPrices()
	items := saveItems(prices)
	saveMonsters(items)

	verifyPrices(items, prices)

	fmt.Println("BYE")
}

func saveItems(prices prices) []*out.Item {
	defer logTime("Items")()

	data, err := ioutil.ReadFile(filepath.Join(dataPath, itemsOtbFile))
	if err != nil {
		panic(err)
	}
	serverClient, clientServer := ReadOtb(data)

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
			attrItem := item
			if equipTo := item.Attributes.Read("transformEquipTo"); equipTo != 0 {
				for _, it := range input {
					if it.Id == int(equipTo) {
						attrItem = it
						break
					}
				}
			}
			it := out.NewItem(client, item, attrItem)
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

	if _, err := fmt.Fprintf(f, `// Code generated by "gobiassets" using 'go generate'. DO NOT EDIT.

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
`, packageName, out.ItemCategoryPrefix(), out.ItemHeader(), varStr.String(), categoriesStr.String(), switchStr.String()); err != nil {
		panic(err)
	}

	verifyClientId(items, input, clientServer)
	return items
}

func saveMonsters(items []*out.Item) []*out.Monster {
	defer logTime("Monsters")()

	f, err := os.OpenFile(monstersFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	monsterFile := readMonsterFile()

	itemsByServerId := make(map[uint16]*out.Item)
	for _, item := range items {
		itemsByServerId[item.ServerId] = item
	}
	monsters := make(map[out.MonsterCategory][]*out.Monster)
	variables := make(map[*out.Monster]string)
	var variableChars int
	for cat, dir := range map[out.MonsterCategory]string{
		out.MonsterCategoryMonsters: "exp",
		out.MonsterCategoryBosses:   "bosses",
		out.MonsterCategorySaga:     "saga",
	} {
		input := readMonsters(dir)
		mon := make([]*out.Monster, len(input))
		for i, monster := range input {
			meta := monsterFile[monster.Path]
			m := out.GetMonster(meta.Id, monster, itemsByServerId)
			mon[i] = m
			variables[m] = meta.Name
			varName := Variable(meta.Name)
			if len(varName) > variableChars {
				variableChars = len(varName)
			}
		}
		sort.Sort(monsterByLevel(mon))
		monsters[cat] = mon
	}

	variableChars++

	var ret []*out.Monster

	// categories
	var categoriesStr strings.Builder
	for cat := out.MonsterCategoryFirst; cat <= out.MonsterCategoryLast; cat++ {
		categoriesStr.WriteString(fmt.Sprintf("\tMonsterCategory(%d): {", cat))
		if monsters, ok := monsters[cat]; ok {
			if len(monsters) > 0 {
				categoriesStr.WriteByte('\n')
			}
			for _, monster := range monsters {
				ret = append(ret, monster)
				categoriesStr.WriteString("\t\t")
				categoriesStr.WriteString(Variable(variables[monster]))
				categoriesStr.WriteString(",\n")
			}
			if len(monsters) > 0 {
				categoriesStr.WriteByte('\t')
			}
		}
		categoriesStr.WriteString("},\n")
	}

	var varStr strings.Builder
	for _, monster := range ret {
		varName := Variable(variables[monster])

		// var
		varStr.WriteString(fmt.Sprintf("\t%s", varName))
		nameDiff := variableChars - len(varName)
		for i := 0; i < nameDiff; i++ {
			varStr.WriteByte(' ')
		}
		varStr.WriteString(fmt.Sprintf("= %s\n", monster.String()))
	}
	if _, err := fmt.Fprintf(f, `// Code generated by "gobiassets" using 'go generate'. DO NOT EDIT.

package %s

%s

%s

%s

var (
%s)

var MonstersCategory = map[MonsterCategory][]*Monster{
%s}
`, packageName, out.MonsterCategoryPrefix(), out.BossClassPrefix(), out.MonsterHeader(), varStr.String(), categoriesStr.String()); err != nil {
		panic(err)
	}
	return ret
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

func verifyClientId(items []*out.Item, initems []*in.Item, clientServer ClientServerMap) {
	desired := make(map[uint16][]*in.Item)
	for _, item := range items {
		if servers, ok := clientServer[item.ClientId]; ok {
			srv := make([]*in.Item, 0, len(servers))
			for _, serverId := range servers {
				if serverId == item.ServerId {
					continue
				}
				for _, it := range initems {
					if it.Id == int(serverId) {
						if it.Name != "" {
							srv = append(srv, it)
							break
						}
					}
				}
			}
			if len(srv) > 0 {
				desired[item.ServerId] = srv
			}
		}
	}
	if len(desired) == 0 {
		return
	}

	var sb strings.Builder
	sb.WriteString("Items with reused client ids:\n")
	for id, servers := range desired {
		fmt.Fprintf(&sb, "%d:%v, ", id, servers)
	}
	log.Println(sb.String())
}

func readItems() []*in.Item {
	f, err := os.Open(filepath.Join(dataPath, itemsXmlFile))
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

type idName struct {
	Id   uint16
	Name string
}

func readMonsterFile() map[string]idName {
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
	ret := make(map[string]idName)
	for _, monster := range m.Monsters {
		ret[monster.File] = idName{monster.Id, monster.Name}
	}
	return ret
}

func readMonsters(dir string) []*in.Monster {
	basePath := filepath.Join(dataPath, monsterDir, dir)
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		panic(err)
	}

	var monsters []*in.Monster
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
		m := in.NewMonster(path.Join(dir, filepath.Base(fpath)))
		if err := xml.NewDecoder(f).Decode(m); err != nil {
			panic(err)
		}
		monsters = append(monsters, m)
	}
	return monsters
}

type monsterByLevel []*out.Monster

func (s monsterByLevel) Len() int      { return len(s) }
func (s monsterByLevel) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s monsterByLevel) Less(i, j int) bool {
	if s[i].Level == s[j].Level {
		if s[i].Health == s[j].Health {
			if s[i].Experience == s[j].Experience {
				return s[i].Name < s[j].Name
			}
			return s[i].Experience < s[j].Experience
		}
		return s[i].Health < s[j].Health
	}
	return s[i].Level < s[j].Level
}
