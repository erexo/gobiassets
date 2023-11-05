package in

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Items struct {
	Items []*Item `xml:"item"`
}

type Item struct {
	Id         int        `xml:"id,attr"`
	Name       string     `xml:"name,attr"`
	Attributes Attributes `xml:"attribute"`
}

func (i *Item) String() string {
	return fmt.Sprintf("%s(%d)", i.Name, i.Id)
}

type Attributes map[string]string

func (a *Attributes) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	if *a == nil {
		*a = make(Attributes)
	}
	var key, value string
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "key":
			key = strings.ToLower(attr.Value)
		case "value":
			value = attr.Value
		}
	}
	(*a)[key] = value
	dec.Skip()
	return nil
}

func (a Attributes) Read(names ...string) int64 {
	var ret int64
	for _, name := range names {
		if v, ok := a[strings.ToLower(name)]; ok {
			value, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Println(err)
			}
			ret += value
		}
	}
	return ret
}

func (a Attributes) ReadString(name string) string {
	if v, ok := a[strings.ToLower(name)]; ok {
		return v
	}
	return ""
}

func (a Attributes) ReadPercent(name string) int64 {
	v := a.Read(name)
	if v == 0 {
		return 0
	}
	return v - 100
}
