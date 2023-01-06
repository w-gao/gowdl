package internal

import (
	"encoding/json"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/internal/domain"
	"github.com/w-gao/gowdl/parsers"
)

type WdlBuilder struct {
	Url            string
	MaxImportDepth int

	Version  string
	Document *domain.Document
}

func NewWdlBuilder(url string) (*WdlBuilder, error) {
	version, err := GetVersion(url)
	if err != nil {
		return nil, err
	}

	if !IsIn(version, []string{"1.0", "1.1", "development"}) {
		return nil, fmt.Errorf("Unsupported WDL version: %s", version)
	}

	maxImportDepth := 10
	builder := &WdlBuilder{Url: url, MaxImportDepth: maxImportDepth, Version: version}
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

	// v1.1
	// input := antlr.NewInputStream(data)
	// lexer := v1_1.NewWdlV1_1Lexer(input)
	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// p := v1_1.NewWdlV1_1Parser(stream)

	tree := p.Document()
	visitor := NewWdlVisitor(this.Version)
	document := visitor.Visit(tree)

	fmt.Printf("%v\n", document)

	out, err := json.MarshalIndent(document, "", "    ")
	if err == nil {
		fmt.Printf("%v\n", string(out))
	}

	return nil
}
