package lexer

import (
	"bufio"
	"strings"
	"unicode"

	"github.com/JairAntonio22/pkg/structs"
)

type Lexer struct {
	reader    *bufio.Reader
	currRune  rune
	currToken structs.Token
	line      uint
	column    uint
	builder   strings.Builder
}

func NewLexer(r *bufio.Reader) Lexer {
	return Lexer{reader: r}
}

func (l Lexer) Peek() structs.Token {
	return l.currToken
}

func (l *Lexer) Read() structs.Token {
	r := l.readRune()

	if unicode.IsLetter(r) {
		l.readKeywordOrId()
		return l.currToken
	}

	return structs.Token{}
}

func (l *Lexer) readRune() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		panic(err) // when does ReadRune returns error?
	}

	l.column += 1

	if r == '\n' {
		l.line += 1
		l.column = 0
	}

	l.currRune = r
	return r
}

func (l *Lexer) readKeywordOrId() {
	column := l.column
	l.builder.WriteRune(l.currRune)

	for {
		r := l.readRune()

		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			break
		}

		l.builder.WriteRune(r)
	}

	l.buildCurrToken(column)
}

func (l *Lexer) buildCurrToken(column uint) {
	l.currToken.Line = l.line
	l.currToken.Column = column

	literal := l.builder.String()
	l.builder.Reset()

	tokType, exists := structs.MapLiteralTokType[literal]

	if exists {
		l.currToken.TokType = tokType
		return
	}

	l.currToken.TokType = structs.TokTypeId
	l.currToken.Literal = &literal
}
