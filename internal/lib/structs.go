package lib

type Document struct {
	Version          string
	Workflow         *Workflow // optional
	DocumentElements []DocumentElement
}

type Workflow struct {
}

type DocumentElement interface{}

type Import struct {
	DocumentElement

	Url   string
	Alias string
}
