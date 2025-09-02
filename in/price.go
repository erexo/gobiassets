package in

import (
	"encoding/xml"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
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

	luaPrices, err := parsePricesFile(filepath.Join(dataPath, npcDir, npcPricesFile))
	if err != nil {
		panic(err)
	}

	for k, v := range luaPrices {
		if cv, ok := ret[k]; ok && cv >= v {
			//panic(fmt.Sprintf("lua price for %d [%d] already contain lua price [%d]", key, cv, value))
			continue
		}
		ret[k] = v
	}

	return ret
}

func parsePricesFile(filePath string) (map[int]int, error) {
	// Read file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := string(data)

	// Regex to match getItems(firstParam, secondParam)
	getItemsRe := regexp.MustCompile(`getItems\s*\(\s*(?:nil|\{[^{}]*\})\s*,\s*(nil|\{[^{}]*\})\s*\)`)
	secondParamMatches := getItemsRe.FindAllStringSubmatch(content, -1)

	result := make(map[int]int)

	// Regex for key-value pairs: [123] = 456,
	pairRe := regexp.MustCompile(`\[(\d+)\]\s*=\s*(\d+)\s*,?`)

	for _, match := range secondParamMatches {
		if len(match) < 2 {
			continue
		}

		for _, param := range []string{match[0], match[1]} {
			param = strings.TrimLeft(param, "getItems(")

			// Skip if second param is nil
			if param == "nil" {
				continue
			}

			// Extract key-value pairs from secondParam
			pairs := pairRe.FindAllStringSubmatch(param, -1)
			for _, p := range pairs {
				if len(p) != 3 {
					continue
				}
				key, err := strconv.Atoi(p[1])
				if err != nil {
					return nil, err
				}
				value, err := strconv.Atoi(p[2])
				if err != nil {
					return nil, err
				}

				if cv, ok := result[key]; ok && cv >= value {
					//panic(fmt.Sprintf("lua price for %d [%d] already contain lua price [%d]", key, cv, value))
					continue
				}

				result[key] = value
			}
		}
	}

	return result, nil
}
