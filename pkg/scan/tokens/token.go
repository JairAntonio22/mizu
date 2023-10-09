package tokens

type Token struct {
	TokType TokType
	Literal string
}

func NewToken(tokType TokType, literal string) Token {
	token, exists := LiteralTokenMap[literal]

	if exists {
		return token
	}

	return Token{
		TokType: tokType,
		Literal: literal,
	}
}

func (t Token) String() string {
	return t.Literal
}

var (
	Eof       = Token{TokType: TypeEof, Literal: "<EOF>"}
	Eol       = Token{TokType: TypeEof, Literal: "<EOL>"}
	Comma     = Token{TokType: TypeComma, Literal: ","}
	Colon     = Token{TokType: TypeColon, Literal: ":"}
	Semicolon = Token{TokType: TypeSemicolon, Literal: ";"}
	LParen    = Token{TokType: TypeLParen, Literal: "("}
	RParen    = Token{TokType: TypeRParen, Literal: ")"}
	LBrace    = Token{TokType: TypeLBrace, Literal: "{"}
	RBrace    = Token{TokType: TypeRBrace, Literal: "}"}
	LBracket  = Token{TokType: TypeLBracket, Literal: "["}
	RBracket  = Token{TokType: TypeRBracket, Literal: "]"}
	DefConst  = Token{TokType: TypeDefConst, Literal: "::"}
	DefVar    = Token{TokType: TypeDefVar, Literal: ":="}
	Assign    = Token{TokType: TypeAssign, Literal: "="}
	Add       = Token{TokType: TypeAdd, Literal: "+"}
	Sub       = Token{TokType: TypeSub, Literal: "-"}
	Mul       = Token{TokType: TypeMul, Literal: "*"}
	Div       = Token{TokType: TypeDiv, Literal: "/"}
	Mod       = Token{TokType: TypeMod, Literal: "%"}
	Eq        = Token{TokType: TypeEq, Literal: "=="}
	Neq       = Token{TokType: TypeNeq, Literal: "!="}
	Less      = Token{TokType: TypeLess, Literal: "<"}
	Great     = Token{TokType: TypeGreat, Literal: ">"}
	Leq       = Token{TokType: TypeLeq, Literal: "<="}
	Geq       = Token{TokType: TypeGeq, Literal: ">="}
	Dot       = Token{TokType: TypeDot, Literal: "."}
	Question  = Token{TokType: TypeQuestion, Literal: "?"}
	If        = Token{TokType: TypeIf, Literal: "if"}
	Else      = Token{TokType: TypeElse, Literal: "else"}
	Switch    = Token{TokType: TypeSwitch, Literal: "switch"}
	Case      = Token{TokType: TypeCase, Literal: "case"}
	Loop      = Token{TokType: TypeLoop, Literal: "loop"}
	Skip      = Token{TokType: TypeSkip, Literal: "skip"}
	Break     = Token{TokType: TypeBreak, Literal: "break"}
	Return    = Token{TokType: TypeReturn, Literal: "return"}
	Not       = Token{TokType: TypeNot, Literal: "not"}
	And       = Token{TokType: TypeAnd, Literal: "and"}
	Or        = Token{TokType: TypeOr, Literal: "or"}
	True      = Token{TokType: TypeTrue, Literal: "true"}
	False     = Token{TokType: TypeFalse, Literal: "false"}
	Nil       = Token{TokType: TypeNil, Literal: "nil"}
)

//go:generate stringer -type=TokType
type TokType uint

const (
	TypeEof TokType = iota
	TypeEol
	TypeComma
	TypeColon
	TypeSemicolon
	TypeLParen
	TypeRParen
	TypeLBrace
	TypeRBrace
	TypeLBracket
	TypeRBracket
	TypeDefConst
	TypeDefVar
	TypeAssign
	TypeAdd
	TypeSub
	TypeMul
	TypeDiv
	TypeMod
	TypeEq
	TypeNeq
	TypeLess
	TypeGreat
	TypeLeq
	TypeGeq
	TypeDot
	TypeQuestion
	TypeIf
	TypeElse
	TypeSwitch
	TypeCase
	TypeLoop
	TypeSkip
	TypeBreak
	TypeReturn
	TypeNot
	TypeAnd
	TypeOr
	TypeTrue
	TypeFalse
	TypeNil
	TypeIdentifier
	TypeInteger
	TypeFloat
	TypeString
)

var LiteralTokenMap = map[string]Token{
	",":      Comma,
	":":      Colon,
	";":      Semicolon,
	"(":      LParen,
	")":      RParen,
	"{":      LBrace,
	"}":      RBrace,
	"[":      LBracket,
	"]":      RBracket,
	"::":     DefConst,
	":=":     DefVar,
	"=":      Assign,
	"+":      Add,
	"-":      Sub,
	"*":      Mul,
	"/":      Div,
	"%":      Mod,
	"==":     Eq,
	"!=":     Neq,
	"<":      Less,
	">":      Great,
	"<=":     Leq,
	">=":     Geq,
	".":      Dot,
	"?":      Question,
	"if":     If,
	"else":   Else,
	"switch": Switch,
	"case":   Case,
	"loop":   Loop,
	"skip":   Skip,
	"break":  Break,
	"return": Return,
	"not":    Not,
	"and":    And,
	"or":     Or,
	"true":   True,
	"false":  False,
	"nil":    Nil,
}
