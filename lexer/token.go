package lexer

import "fmt"

type TokenKind int

// const token which hold different types
const (
	EOF TokenKind = iota
	AUTHOR
	AFFILIATON
	FIELD
	TEXT
	NUMBER
	SPACE
	LINEBREAK

	L_CURLY
	R_CURLY
	L_BRACKET
	R_BRACKET

	ASSIGN
	AT
	COMMA
	DOT
	PLUS
	HYPHEN

	ACUTE
	GRAVE
	BREVE
	HAT
	TILDE
	CIRUM

	INVALID
)

// Token struct will contains token value and it's type
// e.g.: (hyphen, "-")
type Token struct {
	Kind  TokenKind // token type like number, symbol etc...
	Value string    // token value as string
}

// map will store all avilable field names and it's type,
// to check while generate tokens
var fields = map[string]TokenKind{
	"name":     FIELD,
	"mail":     FIELD,
	"phone":    FIELD,
	"url":      FIELD,
	"affil":    FIELD,
	"note":     FIELD,
	"id":       FIELD,
	"div":      FIELD,
	"org":      FIELD,
	"street":   FIELD,
	"landmark": FIELD,
	"postcode": FIELD,
	"postbox":  FIELD,
	"city":     FIELD,
	"state":    FIELD,
	"country":  FIELD,
}

// create token init for given type and value
func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

// Stringer impl for debug purpose to show type as string
func (t Token) String() string {
	return fmt.Sprintf("(%s, %q)", TokenString(t.Kind), t.Value)
}

// convert token-type(int) to string for debug
func TokenString(t TokenKind) string {
	switch t {
	case EOF:
		return "eof"
	case AUTHOR:
		return "author"
	case AFFILIATON:
		return "affi"
	case FIELD:
		return "field"
	case TEXT:
		return "text"
	case NUMBER:
		return "number"
	case L_CURLY:
		return "group begin"
	case R_CURLY:
		return "group end"
	case L_BRACKET:
		return "braket begin"
	case R_BRACKET:
		return "braket end"
	case ASSIGN:
		return "assign"
	case AT:
		return "at"
	case COMMA:
		return "comma"
	case DOT:
		return "dot"
	case PLUS:
		return "plus"
	case HYPHEN:
		return "hyphen"
	case INVALID:
		return "invalid"
	case SPACE:
		return "space"
	case LINEBREAK:
		return "line break"
	default:
		return "unknown"
	}
}
