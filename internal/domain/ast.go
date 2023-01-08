package domain

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// The Abstract Syntax Tree (AST) of a WDL document. A domain.Document struct
// represents the direct translation of the input WDL document into the AST.
// There is no topological ordering of the dependency graph at this stage yet.

// An identifier should match pattern /[a-zA-Z][a-zA-Z0-9_]+/
// See: https://github.com/openwdl/wdl/blob/main/versions/1.0/SPEC.md#whitespace-strings-identifiers-constants
type Identifier string

func (this Identifier) IsValid() bool {
	match, err := regexp.MatchString(`[a-zA-Z][a-zA-Z0-9_]+`, string(this))
	if err != nil {
		return false
	}

	return match
}

type Document struct {
	// This is not part of WDL, but is useful to keep track of.
	Url string `json:"url"`

	Version  string    `json:"version"`
	Workflow *Workflow `json:"workflow,omitempty"` // optional

	Imports []Import `json:"imports,omitempty"` // optional
	Structs []Struct `json:"structs,omitempty"` // optional
	Tasks   []Task   `json:"tasks,omitempty"`   // optional
}

func (this Document) String() string {
	out, err := json.Marshal(this)
	if err != nil {
		return "Document<failed to serialize>"
	}

	return fmt.Sprintf("Document<%s>", out)
}

type Workflow struct {
	Name   Identifier    `json:"name"`
	Inputs []Declaration `json:"inputs,omitempty"`
}

type Import struct {
	Url         string                    `json:"url"`
	AbsoluteUrl string                    `json:"absoluteUrl,omitempty"` // not part of WDL, but useful when resolving imports
	As          Identifier                `json:"as,omitempty"`          // optional
	Aliases     map[Identifier]Identifier `json:"aliases,omitempty"`     // optional
}

type Task struct {
}

type Struct struct {
}

type Type struct {
	Optional bool `json:"optional,omitempty"`
}

type ArrayType struct {
}

// Declaration can be bounded (assigned to expression) or unbounded (no assignment).
type Declaration struct {
	Type       Type        `json:"type"`
	Identifier Identifier  `json:"identifier"`
	Expr       *Expression `json:"expr,omitempty"`
}

type Any interface{}

type IExpression interface {
	// Eval should exist for the parsed runtime DAG, not here.
	// Eval(map[string]Any) Any
}

type Expression struct {
	test string
}
