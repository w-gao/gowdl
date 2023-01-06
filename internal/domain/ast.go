package domain

import (
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

	Imports []Import `json:"imports,omitempty"`
	Structs []Struct `json:"structs,omitempty"`
	Tasks   []Task   `json:"tasks,omitempty"`
}

func (this Document) String() string {
	return fmt.Sprintf("Document<version=%s,imports=%q,structs=%q,tasks=%q>", this.Version, this.Imports, this.Structs, this.Tasks)
}

type Workflow struct {
}

type Import struct {
	Url string     `json:"url"`
	As  Identifier `json:"as,omitempty"`

	// aliases for WDL structs
	Aliases map[string]string `json:"aliases,omitempty"`
}

func (this Import) String() string {
	return fmt.Sprintf("Import<url=%s, as=%s>", this.Url, this.As)
}

type Task struct {
}

type Struct struct {
}

type Expression interface{}

type String struct {
}
