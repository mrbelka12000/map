package models

import "encoding/xml"

type Oms struct {
	XMLName     xml.Name   `xml:"osm"`
	Text        string     `xml:",chardata"`
	Version     string     `xml:"version,attr"`
	Generator   string     `xml:"generator,attr"`
	Copyright   string     `xml:"copyright,attr"`
	Attribution string     `xml:"attribution,attr"`
	License     string     `xml:"license,attr"`
	Bounds      Bounds     `xml:"bounds"`
	Node        []Node     `xml:"node"`
	Way         []Way      `xml:"way"`
	Relation    []Relation `xml:"relation"`
}

type Bounds struct {
	Minlat string `xml:"minlat,attr"`
	Minlon string `xml:"minlon,attr"`
	Maxlat string `xml:"maxlat,attr"`
	Maxlon string `xml:"maxlon,attr"`
}
type Node struct {
	ID        string  `xml:"id,attr"`
	Visible   string  `xml:"visible,attr"`
	Version   *string `xml:"-"`
	Changeset *string `xml:"-"`
	Timestamp *string `xml:"-"`
	User      *string `xml:"-"`
	Uid       *string `xml:"-"`
	Lat       string  `xml:"lat,attr"`
	Lon       string  `xml:"lon,attr"`
	Tag       []Tag   `xml:"tag"`
	MapTag    map[string]string
}

type Way struct {
	ID        string  `xml:"id,attr"`
	Visible   string  `xml:"visible,attr"`
	Version   *string `xml:"-"`
	Changeset *string `xml:"-"`
	Timestamp *string `xml:"-"`
	User      *string `xml:"-"`
	Uid       *string `xml:"-"`
	Nd        []struct {
		Ref string `xml:"ref,attr"`
	} `xml:"nd"`
	Tag []Tag `xml:"tag"`
}

type Relation struct {
	ID        string  `xml:"id,attr"`
	Visible   string  `xml:"visible,attr"`
	Version   string  `xml:"version,attr"`
	Changeset *string `xml:"-"`
	Timestamp *string `xml:"-"`
	User      *string `xml:"-"`
	Uid       *string `xml:"-"`
	Member    []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Ref  string `xml:"ref,attr"`
		Role string `xml:"role,attr"`
	} `xml:"member"`
	Tag []Tag `xml:"tag"`
}

type Tag struct {
	Text string `xml:",chardata"`
	K    string `xml:"k,attr"`
	V    string `xml:"v,attr"`
}
