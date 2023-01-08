package internal

import (
	"fmt"

	"github.com/w-gao/gowdl/internal/domain"
	"github.com/w-gao/gowdl/parsers"
)

type WdlBuilder struct {
	Url      string
	Version  string
	Document *domain.Document
}

func NewWdlBuilder(url string) (*WdlBuilder, error) {
	version, err := GetVersion(url)
	if err != nil {
		return nil, err
	}

	if !IsIn(version, []string{"1.0", "1.1", "development"}) {
		return nil, fmt.Errorf("Unrecognized WDL version: %s", version)
	}

	builder := &WdlBuilder{Url: url, Version: version}
	return builder, nil
}

func (this *WdlBuilder) ParseDocument() (*domain.Document, error) {
	data, err := ReadString(this.Url)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from URI: %w", err)
	}

	visitor := NewWdlVisitor(this.Url, this.Version)
	errListener := &CustomErrorListener{}
	var tree domain.IDocumentContext

	switch this.Version {
	case "1.0":
		p := parsers.NewWdlV1_0Parser(data)
		p.RemoveErrorListeners()
		p.AddErrorListener(errListener)
		tree = p.Document()
	case "1.1":
		p := parsers.NewWdlV1_1Parser(data)
		p.RemoveErrorListeners()
		p.AddErrorListener(errListener)
		tree = p.Document()
	case "development":
		return nil, fmt.Errorf("The development version is not supported yet")
	default:
		return nil, fmt.Errorf("Unknown WDL version: %v", this.Version)
	}

	if len(errListener.Errors) > 0 {
		for _, err := range errListener.Errors {
			fmt.Printf("%v\n", err)
		}

		return nil, fmt.Errorf("SyntaxError detected. Please fix your WDL.")
	}

	document := visitor.VisitDocument(tree)

	return document, nil
}
