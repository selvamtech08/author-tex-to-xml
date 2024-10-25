package ast

import "encoding/xml"

// Affiliation information will be stored
type AffilStmt struct {
	XMLName xml.Name `xml:"affil"`
	ID      int      `xml:"id,attr"`
	OrgDiv  string   `xml:"org-div"`
	OrgName string   `xml:"org-name"`
	Address Address  `xml:"address"`
}

// Address information will be stored
type Address struct {
	Street   string `xml:"street,omitempty"`
	Landmark string `xml:"landmark,omitempty"`
	Pincode  string `xml:"pincode,omitempty"`
	Postbox  string `xml:"postbox,omitempty"`
	City     string `xml:"city,omitempty"`
	State    string `xml:"state,omitempty"`
	Country  string `xml:"country,omitempty"`
}

func (a AffilStmt) stmt() {}

// Will hold Author and Affiliation informations
type AuthorGroupStmt struct {
	XMLName  xml.Name `xml:"author-group"`
	Comment  string   `xml:",comment"`
	Elements []Stmt
}
