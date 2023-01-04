package internal

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/parsers"
)

type WdlBuilder struct {
	Version string
	Url     string

	Document *Document
}

func NewWdlBuilder(url string) (*WdlBuilder, error) {
	version, err := GetVersion(url)
	if err != nil {
		return nil, err
	}

	if !IsIn(version, []string{"1.0", "1.1", "development"}) {
		return nil, fmt.Errorf("Unsupported WDL version: %s", version)
	}

	builder := &WdlBuilder{Version: version, Url: url}
	return builder, nil
}

func (this *WdlBuilder) ParseDocument() error {
	data, err := ReadString(this.Url)
	if err != nil {
		return err
	}

	input := antlr.NewInputStream(data)
	lexer := parsers.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsers.NewWdlV1Parser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	tree := p.Document()
	visitor := NewWdlV1Visitor()
	document := visitor.VisitDocument(tree.(*parsers.DocumentContext))

	fmt.Printf("%v\n", document)

	return nil
}
