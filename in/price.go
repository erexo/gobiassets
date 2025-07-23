package in

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Prices map[int]int

func ReadPrices() Prices {
	ret := make(Prices)
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
							if strings.TrimSpace(values[0]) != "" {
								panic(fmt.Sprintf("Invalid values parameters for '%s' in %s", item, path))
							}
							continue
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
