package out

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"

	"github.com/erexo/gobiassets/in"
)

var disabledNPCs = []string{"Anbu"}

func SaveMapNPCs() {
	spawns := in.ReadSpawnsFile()

	npcs := in.ReadNPCs()

	scmap, _ := in.ReadOtb()

	_ = os.MkdirAll(outputDir, 0755)
	f, err := os.OpenFile(path.Join(outputDir, npcMapFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, npc := range spawns.NPCs {
		npcData, ok := npcs[npc.Name]

		titleNpc := Title(npcData.Name)
		if slices.Contains(disabledNPCs, titleNpc) {
			continue
		}

		if !ok {
			fmt.Println("NOT FOUND", npc.Name)
			continue
		}
		l := npcData.Look
		t, look := "looktype", l.Type
		if l.Typeex != "" {
			id, err := strconv.Atoi(l.Typeex)
			if err != nil {
				panic("failed to decode .." + err.Error())
			}
			cid, ok := scmap[uint16(id)]
			if !ok {
				panic("failed to find id .." + strconv.Itoa(id))
			}

			t = "item"
			look = strconv.Itoa(int(cid))
		}
		if l.Head != 0 || l.Body != 0 || l.Legs != 0 || l.Feet != 0 || l.Addons != 0 {
			t = "looktypeEx"
			look = fmt.Sprintf("%v, %d, %d, %d, %d, %d", look, l.Head, l.Body, l.Legs, l.Feet, l.Addons)
		}

		fmt.Fprintf(f, "%s(\"%s\", %v, atlas.Pos(%d, %d, %d))\n", t, titleNpc, look, npc.X, npc.Y, npc.Z)
	}
}
