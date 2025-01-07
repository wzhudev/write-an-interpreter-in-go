package lexer

import (
	"monkey/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	Col          int
	Row          int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, Row: 1, Col: 0}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	if l.ch == '\n' {
		l.Row += 1
		l.Col = 1
	} else {
		l.Col += 1
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			col := l.Col
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "==", Pos: &token.Position{Row: l.Row, Col: col}}
		} else {
			tok = newToken(token.ASSIGN, l)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l)
	case '(':
		tok = newToken(token.LPAREN, l)
	case ')':
		tok = newToken(token.RPAREN, l)
	case '{':
		tok = newToken(token.LBRACE, l)
	case '}':
		tok = newToken(token.RBRACE, l)
	case ',':
		tok = newToken(token.COMMA, l)
	case '+':
		tok = newToken(token.PLUS, l)
	case '-':
		tok = newToken(token.MINUS, l)
	case '/':
		tok = newToken(token.SLASH, l)
	case '!':
		if l.peekChar() == '=' {
			col := l.Col
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!=", Pos: &token.Position{Row: l.Row, Col: col}}
		} else {
			tok = newToken(token.BANG, l)
		}
	case '*':
		tok = newToken(token.ASTERISK, l)
	case '<':
		tok = newToken(token.LT, l)
	case '>':
		tok = newToken(token.GT, l)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Pos = &token.Position{Row: l.Row, Col: l.Col}
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Pos = &token.Position{Row: l.Row, Col: l.Col}
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, l *Lexer) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(l.ch),
		Pos:     &token.Position{Row: l.Row, Col: l.Col},
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
