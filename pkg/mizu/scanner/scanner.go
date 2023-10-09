package scanner

import (
	"bufio"
	"io"
	"strings"
	. "unicode"

	. "github.com/JairAntonio22/pkg/mizu/token"
)

type Scanner struct {
	eof     bool
	reader  *bufio.Reader
	builder strings.Builder
}

func New(r *bufio.Reader) *Scanner {
	return &Scanner{reader: r}
}

func (s *Scanner) Read() Token {
	if s.eof {
		return Eof
	}

	r := s.readRune()

	for IsSpace(r) {
		if r == '\n' {
			return Eol
		}

		r = s.readRune()
	}

	if s.eof {
		return Eof
	}

	if IsLetter(r) {
		return s.readIdentifier(r)
	}

	if IsDigit(r) {
		return s.readNumber(r)
	}

	switch r {
	case ':':
		switch nextRune := s.readRune(); nextRune {
		case ':':
			return DefConst

		case '=':
			return DefVar
		}

		s.reader.UnreadRune()
		return Colon

	case '=':
		if nextRune := s.readRune(); nextRune == '=' {
			return Eq
		}

		s.reader.UnreadRune()
		return Assign

	case '<':
		if nextRune := s.readRune(); nextRune == '=' {
			return Leq
		}

		return Less

	case '>':
		if nextRune := s.readRune(); nextRune == '=' {
			return Geq
		}

		return Great

	case '!':
		if nextRune := s.readRune(); nextRune == '=' {
			return Geq
		}

		s.reader.UnreadRune()
	}

	token, exists := LiteralTokenMap[string(r)]

	if exists {
		return token
	}

	s.eof = true
	return Eof
}

func (s *Scanner) readIdentifier(first rune) Token {
	s.builder.Reset()

	for r := first; IsLetter(r) || IsDigit(r); r = s.readRune() {
		s.builder.WriteRune(r)
	}

	s.reader.UnreadRune()
	return NewToken(TypeIdentifier, s.builder.String())
}

func (s *Scanner) readNumber(first rune) Token {
	s.builder.Reset()

	for r := first; IsDigit(r); r = s.readRune() {
		s.builder.WriteRune(r)
	}

	s.reader.UnreadRune()
	return NewToken(TypeInteger, s.builder.String())
}

func (l *Scanner) readRune() rune {
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
