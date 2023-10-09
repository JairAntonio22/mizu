package scan

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	. "github.com/JairAntonio22/pkg/scan/tokens"
)

type Scanner struct {
	eof     bool
	reader  *bufio.Reader
	builder strings.Builder
}

func NewScanner(r *bufio.Reader) *Scanner {
	return &Scanner{reader: r}
}

func (s *Scanner) Read() Token {
	if s.eof {
		return Eof
	}

	r := s.readRune()

	for unicode.IsSpace(r) {
		if r == '\n' {
			return Eol
		}

		r = s.readRune()
	}

	if s.eof {
		return Eof
	}

	if unicode.IsLetter(r) {
		return s.readIdentifier(r)
	}

	if unicode.IsDigit(r) {
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
	s.builder.WriteRune(first)
	r := s.readRune()

	for unicode.IsLetter(r) || unicode.IsDigit(r) {
		s.builder.WriteRune(r)
		r = s.readRune()
	}

	literal := s.builder.String()
	token := NewToken(TypeIdentifier, literal)
	s.builder.Reset()
	s.reader.UnreadRune()
	return token
}

func (s *Scanner) readNumber(first rune) Token {
	s.builder.WriteRune(first)
	r := s.readRune()

	for unicode.IsDigit(r) {
		s.builder.WriteRune(r)
		r = s.readRune()
	}

	literal := s.builder.String()
	token := NewToken(TypeInteger, literal)
	s.builder.Reset()
	s.reader.UnreadRune()
	return token
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
