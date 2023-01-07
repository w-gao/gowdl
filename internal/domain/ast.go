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
}

type Import struct {
	Url     string                    `json:"url"`
	As      Identifier                `json:"as,omitempty"`      // optional
	Aliases map[Identifier]Identifier `json:"aliases,omitempty"` // optional
}

type Task struct {
}

type Struct struct {
}

type Expression interface{}

type String struct {
}
