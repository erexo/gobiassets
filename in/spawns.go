package in

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

type XMLSpawns struct {
	XMLName xml.Name   `xml:"spawns"`
	Spawns  []XMLSpawn `xml:"spawn"`
}

type XMLSpawn struct {
	CenterX  int         `xml:"centerx,attr"`
	CenterY  int         `xml:"centery,attr"`
	CenterZ  int         `xml:"centerz,attr"`
	Radius   int         `xml:"radius,attr"`
	Monsters []XMLEntity `xml:"monster"`
	NPCs     []XMLEntity `xml:"npc"`
}

type XMLEntity struct {
	Name      string `xml:"name,attr"`
	X         int    `xml:"x,attr"`
	Y         int    `xml:"y,attr"`
	Z         int    `xml:"z,attr"`
	SpawnTime int    `xml:"spawntime,attr"`
}

type Spawns struct {
	Monsters []Spawn
	NPCs     []Spawn
}

type Spawn struct {
	Name      string
	X, Y, Z   int
	SpawnTime int
}

func ReadSpawnsFile() Spawns {
	f, err := os.Open(filepath.Join(dataPath, spawnsXmlFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var s XMLSpawns
	if err := xml.NewDecoder(f).Decode(&s); err != nil {
		panic(err)
	}

	var monsters, npcs []Spawn

	for _, spawn := range s.Spawns {
		for _, monster := range spawn.Monsters {
			monsters = append(monsters, Spawn{
				Name:      monster.Name,
				X:         spawn.CenterX + monster.X,
				Y:         spawn.CenterY + monster.Y,
				Z:         monster.Z,
				SpawnTime: monster.SpawnTime,
			})
		}
		for _, npc := range spawn.NPCs {
			npcs = append(npcs, Spawn{
				Name:      npc.Name,
				X:         spawn.CenterX + npc.X,
				Y:         spawn.CenterY + npc.Y,
				Z:         npc.Z,
				SpawnTime: npc.SpawnTime,
			})
		}
	}

	return Spawns{
		Monsters: monsters,
		NPCs:     npcs,
	}
}
