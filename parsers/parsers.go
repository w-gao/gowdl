package parsers

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/internal"
	"github.com/w-gao/gowdl/parsers/v1_0"
	"github.com/w-gao/gowdl/parsers/v1_1"
)

func NewWdlV1_0Parser(uri string) (*v1_0.WdlV1Parser, error) {
	data, err := internal.ReadString(uri)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from URI: %w", err)
	}

	input := antlr.NewInputStream(data)
	lexer := v1_0.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := v1_0.NewWdlV1Parser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true

	return p, nil
}

func NewWdlV1_1Parser(uri string) (*v1_1.WdlV1_1Parser, error) {
	data, err := internal.ReadString(uri)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from URI: %w", err)
	}

	input := antlr.NewInputStream(data)
	lexer := v1_1.NewWdlV1_1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := v1_1.NewWdlV1_1Parser(stream)

	return p, nil
}
