package parser

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/selvamtech08/author-tex-xml/ast"
	"github.com/selvamtech08/author-tex-xml/lexer"
)

// Parser will contains lexer init and current position to parse
type Parser struct {
	lex        lexer.Lexer // lexer
	currToken  lexer.Token // point to current token from slice tokens
	pos        int         // current token pos
	lineNumber int         // linenumber error tracking
}

// create new parser init
func NewParser(lex lexer.Lexer) Parser {
	p := Parser{lex: lex, pos: 0, lineNumber: 1}
	p.currToken = p.lex.Tokens[p.pos]
	return p
}

// move current and peek token position by 1
func (p *Parser) advance() {
	p.pos++
	if !p.isEof() {
		p.currToken = p.lex.Tokens[p.pos]
	}
}

func (p *Parser) isEof() bool {
	return p.pos >= len(p.lex.Tokens)
}

// parse tokens one by one
func (p *Parser) Parse() {

	var blockStmt ast.AuthorGroupStmt
	blockStmt.Comment = "author and affiliation elements"

	for p.currToken.Kind != lexer.EOF {
		node := p.parseToken()
		blockStmt.Elements = append(blockStmt.Elements, node)
		p.advance()
		p.skipSpaces()
	}

	// for debug purpose
	file, err := os.Create("out.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	srcjson, err := xml.MarshalIndent(blockStmt, "", "  ")
	if err != nil {
		fmt.Println(srcjson)
		return
	}
	file.WriteString(xml.Header)
	file.Write(srcjson)
}
