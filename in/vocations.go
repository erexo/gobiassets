package in

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

type Vocations struct {
	Vocations []*Vocation `xml:"vocation"`
}

type Vocation struct {
	Name    string          `xml:"name,attr"`
	Id      int             `xml:"id,attr"`
	Formula VocationFormula `xml:"formula"`
}

type VocationFormula struct {
	MainRole string `xml:"mainRole,attr"`
	SubRole1 string `xml:"subRole1,attr"`
	SubRole2 string `xml:"subRole2,attr"`
}

func ReadVocationsFile() []*Vocation {
	f, err := os.Open(filepath.Join(dataPath, vocationsXmlFile))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var s Vocations
	if err := xml.NewDecoder(f).Decode(&s); err != nil {
		panic(err)
	}

	return s.Vocations
}
