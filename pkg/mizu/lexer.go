package mizu

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	. "github.com/JairAntonio22/pkg/mizu/tokens"
)

type Lexer struct {
	eof     bool
	reader  *bufio.Reader
	builder strings.Builder
}

func NewLexer(r *bufio.Reader) Lexer {
	return Lexer{reader: r}
}

func (l *Lexer) Read() Token {
	if l.eof {
		return Eof
	}

	r := l.readRune()

	for unicode.IsSpace(r) {
		if r == '\n' {
			return Eol
		}

		r = l.readRune()
	}

	if l.eof {
		return Eof
	}

	if unicode.IsLetter(r) {
		return l.readIdentifier(r)
	}

	if unicode.IsDigit(r) {
		return l.readNumber(r)
	}

	switch r {
	case ':':
		switch nextRune := l.readRune(); nextRune {
		case ':':
			return DefConst

		case '=':
			return DefVar
		}

		l.reader.UnreadRune()
		return Colon

	case '=':
		if nextRune := l.readRune(); nextRune == '=' {
			return Eq
		}

		l.reader.UnreadRune()

	case '<':
		if nextRune := l.readRune(); nextRune == '=' {
			return Leq
		}

		return Less

	case '>':
		if nextRune := l.readRune(); nextRune == '=' {
			return Geq
		}

		return Great

	case '!':
		if nextRune := l.readRune(); nextRune == '=' {
			return Geq
		}

		l.reader.UnreadRune()
	}

	token, exists := LiteralTokenMap[string(r)]

	if exists {
		return token
	}

	l.eof = true
	return Eof
}

func (l *Lexer) readIdentifier(first rune) Token {
	l.builder.WriteRune(first)
	r := l.readRune()

	for unicode.IsLetter(r) || unicode.IsDigit(r) {
		l.builder.WriteRune(r)
		r = l.readRune()
	}

	literal := l.builder.String()
	token := NewToken(TypeIdentifier, literal)
	l.builder.Reset()
	l.reader.UnreadRune()
	return token
}

func (l *Lexer) readNumber(first rune) Token {
	l.builder.WriteRune(first)
	r := l.readRune()

	for unicode.IsDigit(r) {
		l.builder.WriteRune(r)
		r = l.readRune()
	}

	literal := l.builder.String()
	token := NewToken(TypeInteger, literal)
	l.builder.Reset()
	l.reader.UnreadRune()
	return token
}

func (l *Lexer) readRune() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		if err == io.EOF {
			l.eof = true
			return r
		}

		panic(err)
	}

	return r
}
