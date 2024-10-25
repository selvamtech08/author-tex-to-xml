package parser

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/selvamtech08/author-tex-xml/ast"
	"github.com/selvamtech08/author-tex-xml/lexer"
)

func (p *Parser) parseAuthor() ast.AuthorStmt {
	var author ast.AuthorStmt
	author.ID = AuthorCnt
	// advance to next char after \author marco
	p.advance()
	// skip any space b/w \author and {
	p.skipSpaces()
	// check if current char is open curly
	p.expect(lexer.L_CURLY)
	// move current position after begin curly
	p.advance()
	p.skipSpaces()
	// traverse until parse all fields
	for p.currToken.Kind != lexer.R_CURLY {
		fieldName, fieldValue := p.parseField()
		p.skipSpaces()
		// update field values
		err := updateAuthorField(&author, fieldName, fieldValue)
		if err != nil {
			log.Fatalf("\nparser:line: %d - %v\n", p.lineNumber, err)
		}
		// skip comma after each field end
		if p.currToken.Kind == lexer.COMMA {
			p.advance()
			p.skipSpaces()
		}
	}
	return author
}

// parse fields data and return it's name and along with values
func (p *Parser) parseField() (string, string) {
	var kind, value string
	// skip space before field name
	p.skipSpaces()
	// check next token type is field or not
	p.expect(lexer.FIELD)
	// set field name at kind's 1st position
	kind = p.currToken.Value
	// move pos ahead
	p.advance()
	// check if token is assignment
	p.skipSpaces()
	p.expect(lexer.ASSIGN)
	p.advance()
	p.skipSpaces()
	// parse field's value
	value = p.parseText()
	p.skipSpaces()
	p.advance()
	return kind, value
}

func (p *Parser) parseText() string {
	// move after begin curly
	buf := &bytes.Buffer{}
	p.advance()
	p.skipSpaces()
	// parse until find close curly
	for p.currToken.Kind != lexer.R_CURLY {
		buf.WriteString(p.currToken.Value)
		p.advance()
	}
	return buf.String()
}

// check the field names and assign value to respective struct field
func updateAuthorField(author *ast.AuthorStmt, name, value string) error {
	switch name {
	case "name":
		author.Author.Name = value
	case "mail":
		author.Contact.Mail = strings.Split(value, ",")
	case "phone":
		author.Contact.Phone = strings.Split(value, ",")
	case "url":
		author.Contact.Url = strings.Split(value, ",")
	case "note":
		author.Note = value
	case "affil":
		id, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("expected int affil id, got %v", err)
		}
		author.AffID = id
	default:
		return fmt.Errorf("unknown field type")
	}
	return nil
}
