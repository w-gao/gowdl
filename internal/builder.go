package internal

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/w-gao/gowdl/internal/listeners"
	"github.com/w-gao/gowdl/parsers"
)

type WdlBuilder struct {
	Version string
	Url     string
}

func NewWdlBuilder(url string) (*WdlBuilder, error) {

	// TODO: implement utils.SmartOpen() and move these there.

	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return nil, fmt.Errorf("TODO: WDL document from the internet is not supported")
	}

	if strings.HasPrefix(url, "file://") {
		url = url[7:]
	}

	// ---

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

func (this *WdlBuilder) ParseDocument() {
	input, _ := antlr.NewFileStream(this.Url)

	lexer := parsers.NewWdlV1Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsers.NewWdlV1Parser(stream)
	// p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	// p.BuildParseTrees = true

	tree := p.Document() // start from the root
	antlr.ParseTreeWalkerDefault.Walk(listeners.NewWdlV1Listener(), tree)
}
