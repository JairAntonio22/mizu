package structs

import "fmt"

type Token struct {
	TokType TokType
	Literal *string
	Line    uint
	Column  uint
}

func (t Token) String() string {
	if t.Literal == nil {
		return fmt.Sprintf(".%s", t.TokType.String())
	}

	return *t.Literal
}

//go:generate stringer -type=TokType
type TokType uint

const (
	TokTypeEof TokType = iota
	TokTypeColon
	TokTypeSemicolon
	TokTypeLParen
	TokTypeRParen
	TokTypeLBrace
	TokTypeRBrace
	TokTypeLBracket
	TokTypeRBracket
	TokTypeDefConst
	TokTypeDefVar
	TokTypeAdd
	TokTypeSub
	TokTypeMul
	TokTypeDiv
	TokTypeMod
	TokTypeLess
	TokTypeGreat
	TokTypeLeq
	TokTypeGeq
	TokTypeDot
	TokTypeQuestion
	TokTypeIf
	TokTypeElse
	TokTypeSwitch
	TokTypeCase
	TokTypeLoop
	TokTypeSkip
	TokTypeBreak
	TokTypeReturn
	TokTypeNot
	TokTypeAnd
	TokTypeOr
	TokTypeTrue
	TokTypeFalse
	TokTypeNil
	TokTypeId
	TokTypeInteger
	TokTypeFloat
	TokTypeString
)

var MapLiteralTokType = map[string]TokType{
	":":      TokTypeColon,
	";":      TokTypeSemicolon,
	"(":      TokTypeLParen,
	")":      TokTypeRParen,
	"{":      TokTypeLBrace,
	"}":      TokTypeRBrace,
	"[":      TokTypeLBracket,
	"]":      TokTypeRBracket,
	"::":     TokTypeDefConst,
	":=":     TokTypeDefVar,
	"+":      TokTypeAdd,
	"-":      TokTypeSub,
	"*":      TokTypeMul,
	"/":      TokTypeDiv,
	"%":      TokTypeMod,
	"<":      TokTypeLess,
	">":      TokTypeGreat,
	"<=":     TokTypeLeq,
	">=":     TokTypeGeq,
	".":      TokTypeDot,
	"?":      TokTypeQuestion,
	"if":     TokTypeIf,
	"else":   TokTypeElse,
	"switch": TokTypeSwitch,
	"case":   TokTypeCase,
	"loop":   TokTypeLoop,
	"skip":   TokTypeSkip,
	"break":  TokTypeBreak,
	"return": TokTypeReturn,
	"not":    TokTypeNot,
	"and":    TokTypeAnd,
	"or":     TokTypeOr,
	"true":   TokTypeTrue,
	"false":  TokTypeFalse,
	"nil":    TokTypeNil,
}
