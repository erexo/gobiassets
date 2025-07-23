package in

import (
	"encoding/xml"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ReadSageTree() *SageTree {
	f, err := os.Open(filepath.Join(dataPath, sageTreeXmlFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tree := &SageTree{}
	if err := xml.NewDecoder(f).Decode(tree); err != nil {
		panic(err)
	}

	return tree
}

type SageTree struct {
	Nodes []*SageNode `xml:"node"`
	Links []*SageLink `xml:"link"`
}

type SageNode struct {
	Id       uint16     `xml:"id,attr"`
	Type     int        `xml:"type,attr"`
	Icon     int        `xml:"icon,attr"`
	Stats    []NodeStat `xml:"stat"`
	Position Point      `xml:"position,attr"`
}

type SageLink struct {
	Aid uint16 `xml:"aId,attr"`
	Bid uint16 `xml:"bId,attr"`
}

type NodeStat struct {
	Value int16 `xml:"value,attr"`
	Icon  int   `xml:"icon,attr"`
}

type Point image.Point

func (p *Point) UnmarshalXMLAttr(attr xml.Attr) error {
	parts := strings.Split(attr.Value, ";")
	if len(parts) != 2 {
		return fmt.Errorf("invalid position format: %s", attr.Value)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	*p = Point{X: x, Y: y}
	return nil
}
