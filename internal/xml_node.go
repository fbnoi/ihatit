package internal

import "encoding/xml"

//------------ SQL xml
type Mapping struct {
	XmlTag  xml.Name `xml:"mapping"`
	Selects []Select `xml:"select"`
	Inserts []Insert `xml:"insert"`
	Updates []Update `xml:"update"`
	Deletes []Delete `xml:"delete"`
}

type Select struct {
	XmlTag       xml.Name `xml:"select"`
	Query        string   `xml:",chardata"`
	ResultMapper string   `xml:"resultMapper,attr"`
	Params       []*Param `xml:"param"`
}

type Insert struct {
	XmlTag xml.Name `xml:"insert"`
	Query  string   `xml:",chardata"`
	Params []*Param `xml:"param"`
	Batch  *Batch   `xml:"batch"`
}

type Update struct {
	XmlTag xml.Name `xml:"update"`
	Query  string   `xml:",chardata"`
	Params []*Param `xml:"param"`
}

type Delete struct {
	XmlTag xml.Name `xml:"delete"`
	Query  string   `xml:",chardata"`
	Params []*Param `xml:"param"`
}

type Param struct {
	XmlTag xml.Name `xml:"param"`
	Name   string   `xml:"name,attr"`
	Type   string   `xml:"type,attr"`
}

type Batch struct {
	XmlTag xml.Name `xml:"batch"`
	Params []*Param `xml:"param"`
}

//----------- Model xml
type Mapper struct {
	XmlTag xml.Name `xml:"mapper"`
	Models []*Model `xml:"model"`
}

type Model struct {
	XmlTag xml.Name `xml:"model"`
	Name   string   `xml:"name,attr"`
	Key    *Key     `xml:"key"`
	Fields []*Field `xml:"field"`
}

type Field struct {
	XmlTag   xml.Name `xml:"field"`
	Name     string   `xml:"name,attr"`
	Type     string   `xml:"type,attr"`
	Mapped   string   `xml:"mapped,attr"`
	Inverted string   `xml:"inverted,attr"`
	Relation string   `xml:"relation,attr"`
}

type Key struct {
	XmlTag    xml.Name `xml:"key"`
	Name      string   `xml:"name,attr"`
	Generator string   `xml:"generator,attr"`
	Type      string   `xml:"type,attr"`
}
