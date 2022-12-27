package in

import (
	"encoding/xml"
)

type Items struct {
	Items []*Item `xml:"item"`
}

type Item struct {
	Id         int        `xml:"id,attr"`
	Name       string     `xml:"name,attr"`
	Attributes Attributes `xml:"attribute"`
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
			key = attr.Value
		case "value":
			value = attr.Value
		}
	}
	(*a)[key] = value
	dec.Skip()
	return nil
}
