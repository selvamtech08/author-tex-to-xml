package parser

import (
	"fmt"
	"log"
	"strconv"

	"github.com/selvamtech08/author-tex-xml/ast"
	"github.com/selvamtech08/author-tex-xml/lexer"
)

func (p *Parser) parseAffil() ast.AffilStmt {
	var affil ast.AffilStmt
	// advance to next char after \affil marco
	p.advance()
	// skip any space b/w \affil and {
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
		err := updateAffilField(&affil, fieldName, fieldValue)
		if err != nil {
			log.Fatalf("\nparser:line: %d - %v\n", p.lineNumber, err)
		}
		// skip comma after each field end
		if p.currToken.Kind == lexer.COMMA {
			p.advance()
			p.skipSpaces()
		}
	}
	return affil
}

func (p *Parser) parseLineBreak() ast.LineBreakStmt {
	var lineBreak ast.LineBreakStmt
	// skip all spaces
	p.skipSpaces()
	return lineBreak
}

// check the field names and assign value to respective struct field
func updateAffilField(affil *ast.AffilStmt, name, value string) error {
	switch name {
	case "id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("affid should valid int, got '%s'", value)
		}
		affil.ID = id
	case "div":
		affil.OrgDiv = value
	case "org":
		affil.OrgName = value
	case "street":
		affil.Address.Street = value
	case "landmark":
		affil.Address.Landmark = value
	case "postcode":
		affil.Address.Pincode = value
	case "postbox":
		affil.Address.Postbox = value
	case "city":
		affil.Address.City = value
	case "state":
		affil.Address.State = value
	case "country":
		affil.Address.Country = value
	default:
		return fmt.Errorf("unknown field type: `%s`", name)
	}
	return nil
}
