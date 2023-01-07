package internal

import (
	"fmt"

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
		return nil, fmt.Errorf("Unrecognized WDL version: %s", version)
	}

	maxImportDepth := 10
	builder := &WdlBuilder{Url: url, MaxImportDepth: maxImportDepth, Version: version}
	return builder, nil
}

func (this *WdlBuilder) ParseDocument() (*domain.Document, error) {
	data, err := ReadString(this.Url)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from URI: %w", err)
	}

	version := this.Version
	visitor := NewWdlVisitor(version)
	var document *domain.Document

	switch version {
	case "1.0":
		p := parsers.NewWdlV1_0Parser(data)
		document = visitor.VisitDocument(p.Document())
	case "1.1":
		p := parsers.NewWdlV1_1Parser(data)
		document = visitor.VisitDocument(p.Document())
	case "development":
		return nil, fmt.Errorf("The development version is not supported yet")
	default:
		return nil, fmt.Errorf("Unknown WDL version: %v", this.Version)
	}

	return document, nil
}
