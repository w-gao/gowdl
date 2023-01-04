package internal

import "fmt"

type Document struct {
	Version  string
	Workflow *Workflow // optional

	Imports []Import
	Structs []Struct
	Tasks   []Task
}

type Workflow struct {
}

type Import struct {
	Url string
	As  string

	// aliases for WDL structs
	Aliases map[string]string
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
