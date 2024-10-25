package ast

import "encoding/xml"

// Author information will be stored
type AuthorStmt struct {
	XMLName xml.Name   `xml:"author"`
	ID      int        `xml:"id,attr"`
	AffID   int        `xml:"aff-id,attr"`
	Author  AuthorName `xml:"author-name"`
	Contact Contact    `xml:"contact"`
	Note    string     `xml:"note,omitempty"`
}

// AuthorName information will be stored
type AuthorName struct {
	Name string `xml:"name"`
}

// Contact information will be stored
type Contact struct {
	Mail  []string `xml:"mail,omitempty"`
	Phone []string `xml:"phone,omitempty"`
	Url   []string `xml:"url,omitempty"`
}

func (f AuthorStmt) stmt() {}

// Line break information will be stored
type LineBreakStmt struct {
	Value string `xml:",comment"`
}

func (f LineBreakStmt) stmt() {}

// Undefined tokens will be stored
// TODO: it should be fixed later,
// to accept different macros and defintions
type UndefinedStmt struct {
	XMLName xml.Name `xml:"error"`
	Line    int      `xml:"line,attr"`
	Type    string   `xml:"type,attr"`
	Value   string   `xml:",chardata"`
}

func (u UndefinedStmt) stmt() {}
