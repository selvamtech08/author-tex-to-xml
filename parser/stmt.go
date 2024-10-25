package parser

import (
	"fmt"
	"log"

	"github.com/selvamtech08/author-tex-xml/ast"
	"github.com/selvamtech08/author-tex-xml/lexer"
)

var AuthorCnt int

func (p *Parser) parseToken() ast.Stmt {
	switch p.currToken.Kind {
	case lexer.AUTHOR:
		AuthorCnt++
		return p.parseAuthor()
	case lexer.LINEBREAK, lexer.SPACE:
		return p.parseLineBreak()
	case lexer.AFFILIATON:
		return p.parseAffil()
	default:
		fmt.Printf("parser:line: %d - undefined token type found: %q\n", p.lineNumber, lexer.TokenString(p.currToken.Kind))
		return p.parseUndefined()
	}
}

// helper function to check token type with expected type
func (p *Parser) expect(kind lexer.TokenKind) {
	if p.currToken.Kind != kind {
		log.Fatalf("parser:line: %d - expected %v, got %v", p.lineNumber, lexer.TokenString(kind), lexer.TokenString(p.currToken.Kind))
	}
}

// skip space characters around field names
func (p *Parser) skipSpaces() {
	for p.currToken.Kind == lexer.LINEBREAK || p.currToken.Kind == lexer.SPACE {
		if p.currToken.Kind == lexer.LINEBREAK {
			p.lineNumber++
		}
		p.advance()
	}
}

func (p *Parser) parseUndefined() ast.UndefinedStmt {
	return ast.UndefinedStmt{Line: p.lineNumber, Value: p.currToken.Value,
		Type: lexer.TokenString(p.currToken.Kind)}
}
