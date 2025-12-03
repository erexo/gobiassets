package in

import (
	"encoding/xml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type NPCs map[string]NPC

type NPC struct {
	Name         string `xml:"name,attr"`
	Script       string `xml:"script,attr"`
	WalkInterval int    `xml:"walkinterval,attr"`
	FloorChange  int    `xml:"floorchange,attr"`

	Health     Health     `xml:"health"`
	Look       Look       `xml:"look"`
	Parameters Parameters `xml:"parameters"`
}

type Parameters struct {
	Params []Parameter `xml:"parameter"`
}

type Parameter struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

func ReadNPCs() NPCs {
	ret := make(NPCs)

	if err := filepath.WalkDir(filepath.Join(dataPath, npcDir), func(path string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() || filepath.Ext(path) != ".xml" {
			return err
		}
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		var npc NPC
		err = xml.NewDecoder(f).Decode(&npc)
		f.Close()
		if err != nil {
			return err
		}

		ret[filepath.Base(strings.TrimSuffix(path, filepath.Ext(path)))] = npc

		return nil
	}); err != nil {
		panic(err)
	}

	return ret
}
