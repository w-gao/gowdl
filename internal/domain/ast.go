package domain

import "fmt"

// The Abstract Syntax Tree (AST) of a WDL document. A domain.Document struct
// represents the direct translation of the input WDL document into the AST.
// There is no topological ordering of the dependency graph at this stage yet.

type Document struct {
	Version  string    `json:"version"`
	Workflow *Workflow `json:"workflow,omitempty"` // optional

	Imports []Import `json:"imports,omitempty"`
	Structs []Struct `json:"structs,omitempty"`
	Tasks   []Task   `json:"tasks,omitempty"`
}

type Workflow struct {
}

type Import struct {
	Url string `json:"url"`
	As  string `json:"as,omitempty"`

	// aliases for WDL structs
	Aliases map[string]string `json:"aliases,omitempty"`
}

type Task struct {
}

type Struct struct {
}

func (this Document) String() string {
	return fmt.Sprintf("Document<version=%s,imports=%q,structs=%q,tasks=%q>", this.Version, this.Imports, this.Structs, this.Tasks)
}

func (this Import) String() string {
	return fmt.Sprintf("Import<url=%s, as=%s>", this.Url, this.As)
}
