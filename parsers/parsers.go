package parsers

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/parsers/v1_0"
	"github.com/w-gao/gowdl/parsers/v1_1"
)

func NewWdlV1_0Parser(data string) *v1_0.WdlV1Parser {
	input := antlr.NewInputStream(data)
	lexer := v1_0.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := v1_0.NewWdlV1Parser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true

	return p
}

func NewWdlV1_1Parser(data string) *v1_1.WdlV1_1Parser {
	input := antlr.NewInputStream(data)
	lexer := v1_1.NewWdlV1_1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := v1_1.NewWdlV1_1Parser(stream)

	return p
}
