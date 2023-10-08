package lexer

import (
	"bufio"
	"io"
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

	if unicode.IsDigit(r) {
		l.readNumber()
		return l.currToken
	}

	return structs.Token{}
}

func (l *Lexer) readKeywordOrId() {
	l.builder.WriteRune(l.currRune)

	column := l.column
	r := l.readRune()

	for unicode.IsLetter(r) || unicode.IsDigit(r) {
		l.builder.WriteRune(r)
		r = l.readRune()
	}

	l.reader.UnreadRune()
	l.buildCurrToken(column)
}

func (l *Lexer) readNumber() {
	l.builder.WriteRune(l.currRune)

	column := l.column
	r := l.readRune()

	for unicode.IsDigit(r) {
		l.builder.WriteRune(r)
		r = l.readRune()
	}

	l.reader.UnreadRune()
	l.buildCurrToken(column)
}

func (l *Lexer) readRune() rune {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		if r == '\n' {
			l.line += 1
			l.column = 0
		}

		l.column += 1

		if !unicode.IsSpace(r) {
			l.currRune = r
			break
		}
	}

	return l.currRune
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
