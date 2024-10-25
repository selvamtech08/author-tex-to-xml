package lexer

import (
	"fmt"
	"os"
	"unicode"
)

// Lexer struct will generated and store tokens as slice for given input strings
type Lexer struct {
	source     []rune  // input from source file
	currPos    int     // point to current char pos
	currChar   rune    // point to current char
	Tokens     []Token // slice will store all the tokens
	lineNumner int     // for error tracking
}

var symbolList = [...]byte{'`', '\'', '~', '.', '=', '^', 'u', 'v', '"'}

// function to create new lexer init
func NewLexer(input string) Lexer {
	l := Lexer{source: []rune(input), currPos: 0, Tokens: make([]Token, 0)}
	l.currChar = l.source[0]
	l.lineNumner++
	return l
}

// check the source traverse reached end
func (l *Lexer) isEof() bool {
	return l.currPos >= len(l.source)
}

// advance current and peek char position by 1
func (l *Lexer) advance() {
	l.currPos++
	if !l.isEof() {
		l.currChar = l.source[l.currPos]
	} else {
		l.currChar = 0
	}
}

// add the current token to tokens list
func (l *Lexer) push(token Token) {
	l.Tokens = append(l.Tokens, token)
}

// check if the given char is alphabet or not
func isAlpha(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

// get marco string from current position to until find non-alpha character
// return marco string and advance the current positions to final match
func (l *Lexer) getMacro() string {
	pos := l.currPos
	for isAlpha(l.currChar) {
		l.advance()
	}
	return l.getChars(pos)
}

// get text string from current position to until find non-alpha character
// return string and advance the current positions to final match
func (l *Lexer) getText() string {
	pos := l.currPos
	for isAlpha(l.currChar) {
		l.advance()
	}
	return l.getChars(pos)
}

// remove the spaces around field assignment and braces
func (l *Lexer) getSpace() string {
	pos := l.currPos
	for l.currChar == ' ' {
		l.advance()
	}
	return l.getChars(pos)
}

// check if the given char is digit or not
func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

// check if the given char is digit or not
func isAccent(char rune) bool {
	for _, val := range symbolList {
		if val == byte(char) {
			return true
		}
	}
	return false
}

// return number of chars in given range to current
func (l *Lexer) getChars(num int) string {
	// one position extra move in the for loop condition,
	// so adjusted here by -1
	defer func() {
		l.currPos--
	}()
	if l.isEof() {
		return string(l.source[num:])
	}
	return string(l.source[num:l.currPos])
}

// get digit string from current position to until find non-alpha character
// return string and advance the current positions to final match
func (l *Lexer) getNumner() string {
	pos := l.currPos
	for isDigit(l.currChar) {
		l.advance()
	}
	return l.getChars(pos)
}

func (l *Lexer) getAccent(char rune) string {
	l.advance()
	accentUnicode := checkAccent(byte(char), l.currChar)
	return string(accentUnicode)
}

// generate tokens until reached at end of input string
func (l *Lexer) Generate() {
	for l.nextToken().Kind != EOF {
		// empty body
	}
}

// trim the space
func (l *Lexer) skipSpace() {
	for l.currChar == ' ' {
		l.advance()
	}
	l.currPos--
}

// generate token for current character
func (l *Lexer) nextToken() Token {
	var token Token

	char := l.currChar
	switch char {
	case '\\':
		l.advance()
		macro := l.getMacro()
		l.skipSpace()
		if macro == "author" {
			token = NewToken(AUTHOR, macro)
		} else if macro == "affil" {
			token = NewToken(AFFILIATON, macro)
		} else if isAccent(l.currChar) {
			l.advance() // TODO: to check why two times called required
			l.advance()
			accent := l.getAccent(l.currChar)
			token = NewToken(TEXT, accent)
		} else if macro == "u" || macro == "v" {
			// TODO: handle brace
			l.skipSpace()
			l.advance()
			accent := l.getAccent([]rune(macro)[0])
			token = NewToken(TEXT, accent)
		} else {
			// TODO: will add later
			fmt.Printf("lexer:line: %d - unexpected marco found: %q\n", l.lineNumner, macro)
			os.Exit(2)
		}
	case '\r':
		l.advance()
		l.lineNumner++
		token = NewToken(LINEBREAK, string(l.currChar))
	case '\n':
		l.lineNumner++
		token = NewToken(LINEBREAK, string(char))
	case ' ':
		space := l.getSpace()
		token = NewToken(SPACE, space)
	case '\t':
		token = NewToken(SPACE, string(char))
	case '{':
		token = NewToken(L_CURLY, string(char))
	case '}':
		token = NewToken(R_CURLY, string(char))
	case '[':
		token = NewToken(L_BRACKET, string(char))
	case '=':
		token = NewToken(ASSIGN, string(char))
	case ']':
		token = NewToken(R_BRACKET, string(char))
	case ',':
		token = NewToken(COMMA, string(char))
	case '@':
		token = NewToken(AT, string(char))
	case '.':
		token = NewToken(DOT, string(char))
	case '+':
		token = NewToken(PLUS, string(char))
	case '-':
		token = NewToken(HYPHEN, string(char))
	case 0:
		token = NewToken(EOF, string(char))
	default:
		if isAlpha(char) {
			text := l.getText()
			if tag, exists := fields[text]; exists {
				token = NewToken(tag, text)
			} else {
				token = NewToken(TEXT, text)
			}
		} else if isDigit(char) {
			number := l.getNumner()
			token = NewToken(NUMBER, number)
		} else if unicode.IsLetter(char) {
			token = NewToken(TEXT, string(l.currChar))
		} else {
			fmt.Printf("lexer: unknown token found: '%c'\n", char)
			os.Exit(2)
		}
	}
	l.push(token)
	l.advance()
	return token
}
